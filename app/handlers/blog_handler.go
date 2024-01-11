package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserBlogs(c echo.Context) error {

	return c.String(http.StatusAccepted, "aman")
}
