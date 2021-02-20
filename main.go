package main

import (
	"gub/database"
	"gub/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Router - global router
// var Router *gin.Engine

func main() {
	baseURL := "/api"
	gin.ForceConsoleColor()

	// migrating the schema to db
	database.DBsetup()
	user.MigrateUser()
	// user.CreateUser()

	// gin setup
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "welcome to go api",
		})
	})

	user.Routes(r, baseURL)

	r.Run(":5000")
}
