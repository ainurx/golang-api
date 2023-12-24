package models

type Blog struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	UserID  int    `json:"userId" gorm:"references:ID"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
