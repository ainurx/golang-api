package handlers

import (
	"encoding/json"
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

	db, err := common.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	hashedPass, err := common.HashPassword(password)

	if err != nil {
		fmt.Println(err.Error())
	}

	user := models.User{
		Username:  username,
		Password:  hashedPass,
		FirstName: firstName,
		LastName:  lastName,
	}

	result := db.Create(&user)

	fmt.Println(result)
	return c.JSON(http.StatusOK, user)
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
	username := c.FormValue("username")
	password := c.FormValue("password")

	if common.IsEmpty(username) {
		return c.String(http.StatusBadRequest, "Username is required.")
	}

	if common.IsEmpty(password) {
		return c.String(http.StatusBadRequest, "Password is required.")
	}

	db, err := common.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	var user models.User
	db.Find(&user, "username = ?", username)

	result, err := json.Marshal(user)

	if err != nil {
		fmt.Println("error parsing to JSON")
		fmt.Println(err.Error())
	}

	n := len(result)
	s := string(result[:n])
	fmt.Println(s)

	var jsonUser models.User

	err = json.Unmarshal([]byte(s), &jsonUser)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c.JSON(http.StatusAccepted, jsonUser)
}
