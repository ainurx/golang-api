package common

import (
	"fmt"
	"log"
	"os"

	// External Dependencies
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	// Internal Dependencies
	"webapi/app/models"
)

func CreateToken(payload models.User) string {
	var (
		t *jwt.Token
	)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := []byte(os.Getenv("SECRET_KEY"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  payload.Username,
		"firstName": payload.FirstName,
		"lastName":  payload.LastName,
	})
	s, err := t.SignedString(key)

	if err != nil {
		log.Println("Error signing JWT:", err)
		return "Internal Server Error"
	}

	return s
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	// var (
	// 	t *jwt.Token
	// )
	err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	key := []byte(os.Getenv("SECRET_KEY"))

	// t = jwt.New(jwt.SigningMethodHS256)
	// s, err := t.SignedString(token)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println("tidaak aman")
		log.Println(err.Error())
	}
	fmt.Println(token)
	return token, err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
