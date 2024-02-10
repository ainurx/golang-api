package routes

import (
	"github.com/labstack/echo/v4"

	"webapi/app/handlers"
	"webapi/app/middleware"
)

func Routes(e *echo.Echo) {
	e.POST("/signin", handlers.SignIn)

	e.POST("/signup", handlers.SignUp)
	e.GET("/user/:id", handlers.FindUser)

	e.GET("/book", middleware.UserMiddleware(handlers.GetUserBlogs, handlers.ErrorHandlers))
}
