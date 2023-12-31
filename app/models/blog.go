package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	UserID    int    `json:"userId" gorm:"references:ID"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
