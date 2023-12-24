package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	// "webapi/app/common"
	"webapi/app/routes"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "GOLANG API")
	})

	// db, err := common.ConnectDB()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("Database connected successfully")
	// }

	// db.Close()
	// fmt.Println(db)
	// common.CreateTable()

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
