package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	// External Dependencies

	"github.com/labstack/echo/v4"

	// Internal Dependencies
	"webapi/app/common"
	"webapi/app/models"
)

func CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")

	if common.IsEmpty(username) {
		return c.String(http.StatusBadRequest, "Username is required.")
	}

	if common.IsEmpty(password) {
		return c.String(http.StatusBadRequest, "Password is required.")
	}

	if common.IsEmpty(firstName) {
		return c.String(http.StatusBadRequest, "First name is required.")
	}

	if common.IsEmpty(lastName) {
		return c.String(http.StatusBadRequest, "Last name is required.")
	}

	result := make(chan models.User)

	go func() {
		result <- models.User{
			Id:        1,
			Username:  username,
			Password:  password,
			FirstName: firstName,
			LastName:  lastName,
		}
	}()

	hasil := <-result

	return c.JSON(http.StatusOK, hasil)
}

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

	result := models.User{
		Id:        intId,
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
	fmt.Print(result)

	token := common.CreateToken(result)

	return c.String(http.StatusOK, token)
}

func SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "No process yet")
}
