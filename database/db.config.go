package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB - bla bla
var DB *gorm.DB

// DBsetup is for basic DB setup
func DBsetup() {
	var err error
	// db configuration
	dsn := "root:@tcp(127.0.0.1:3306)/gub?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Alhamdulillah, DB connected successfully!")
}
