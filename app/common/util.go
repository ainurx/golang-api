package common

import (
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

func ValidateToken(token string) string {
	var (
		t *jwt.Token
	)
	err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// key := []byte(os.Getenv("SECRET_KEY"))

	t = jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString(token)

	if err != nil {
		log.Println(err.Error())
	}

	return s
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
