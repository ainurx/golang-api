package routes

import (
	"github.com/labstack/echo/v4"

	"webapi/app/handlers"
)

func Routes(e *echo.Echo) {
	e.POST("/user/:id", handlers.GetUser)
}
