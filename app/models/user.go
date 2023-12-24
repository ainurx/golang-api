package models

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Blog      []Blog `gorm:"foreignKey:UserID"`
}
