package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"webapi/app/models"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	name := c.QueryParam("name")
	intId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	result := models.User{
		Id:   intId,
		Name: name,
	}

	return c.JSON(http.StatusOK, result)
}
