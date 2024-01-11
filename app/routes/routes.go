package routes

import (
	"github.com/labstack/echo/v4"

	"webapi/app/handlers"
	"webapi/app/middleware"
)

func Routes(e *echo.Echo) {
	e.POST("/signin", handlers.SignIn)

	e.POST("/user", handlers.CreateUser)
	e.POST("/user/:id", handlers.GetUser)

	e.GET("/book", middleware.UserMiddleware(handlers.GetUserBlogs, handlers.ErrorHandlers))
}
