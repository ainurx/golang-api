package common

import (
	"fmt"
	// "webapi/app/common"

	"webapi/app/models"
)

func CreateTable() {
	db, err := ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)

	// db.Model(&models.User{}).Association("Blogs").ForeignKey("user_id")
	// db.Model(&models.User{}).Association("Blogs").ForeignKey("user_id")
}
