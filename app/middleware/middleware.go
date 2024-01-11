package middleware

import (
	"fmt"
	"webapi/app/common"

	"github.com/labstack/echo/v4"
)

func UserMiddleware(next, stop echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header

		token := header.Get("auth-token")
		// c.Request().Header('')
		user, err := common.ValidateToken(token)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(token)
		fmt.Println(user)
		return stop(c)
		// return next(c)
	}
}
