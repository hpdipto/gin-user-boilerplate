package user

import (
	"fmt"
	db "gub/database"
)

// MigrateUser - migrate the user schema to db
func MigrateUser() {
	db.DB.AutoMigrate(&User{})
	fmt.Println("Alhamdulillah User schema migrated successfully!")
}

// CreateUser is for creating an user
func CreateUser() {
	db.DB.Create(&User{
		FirstName: "Haris",
		LastName:  "Dipto",
		Email:     "haris.dipto@gmail.com",
	})
}
