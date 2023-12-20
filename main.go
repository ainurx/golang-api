package main

import (
	"github.com/labstack/echo/v4"

	"webapi/app/routes"
)

func main() {
	e := echo.New()
	routes.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
