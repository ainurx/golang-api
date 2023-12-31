package models

import (
	"time"
)

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Blog      []Blog `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
