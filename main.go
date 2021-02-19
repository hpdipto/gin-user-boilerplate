package main

import (
	"gub/database"
	"gub/user"
)

func main() {
	// migrating the schema to db
	database.DBsetup()
	user.MigrateUser()
	user.CreateUser()

	// gin setup
	// r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"hello": "welcome to go api",
	// 	})
	// })

	// r.Run(":5000")
}
