package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {

	// db configuration
	dsn := "root:@tcp(127.0.0.1:3306)/gub?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

}