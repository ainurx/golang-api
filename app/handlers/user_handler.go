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

func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "No process yet")
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

	fmt.Println("==== HM ====")

	return c.String(http.StatusOK, token)
}

func SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "No process yet")
}
