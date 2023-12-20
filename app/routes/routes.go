package routes

import (
	"github.com/labstack/echo/v4"

	"webapi/app/handlers"
)

func Routes(e *echo.Echo) {
	e.GET("/user/:id", handlers.GetUser)
}
