package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"webapi/app/routes"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "GOLANG API")
	})

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
