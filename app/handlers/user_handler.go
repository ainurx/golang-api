package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// External Dependencies
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	// Internal Dependencies
	"webapi/app/models"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")

	var (
		t *jwt.Token
		s string
	)
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	key := []byte(os.Getenv("SECRET_KEY"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"firstName": firstName,
		"lastName":  lastName,
	})
	s, err = t.SignedString(key)

	if err != nil {
		log.Println("Error signing JWT:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	result := models.User{
		Id:        intId,
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
	fmt.Print(result)

	return c.String(http.StatusOK, s)
}
