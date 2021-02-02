package user

import (
	"time"

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	gender    string `json:"gender"`
	dob       time.Time `json:"dob"`
	email string `json:"email"`
	phone string `json:"email"`
	password string `json:"password"`
}