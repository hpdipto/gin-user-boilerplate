package database

import (
	"fmt"
	user "gub/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBsetup is for basic DB setup
func DBsetup() *gorm.DB {
	// db configuration
	dsn := "root:@tcp(127.0.0.1:3306)/gub?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Alhamdulillah DB connected successfully!")

	return db
}

// Migrate for migrating all schema to db
func Migrate() {
	db := DBsetup()

	db.AutoMigrate(&user.User{})

	fmt.Println("Alhamdulillah migration done successfully!")
}
