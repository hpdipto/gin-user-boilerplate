package user

import (
	"fmt"
	db "gub/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MigrateUser - migrate the user schema to db
func MigrateUser() {
	db.DB.AutoMigrate(&User{})
	fmt.Println("Alhamdulillah User schema migrated successfully!")
}

// GetUser for getting information about a user
func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user User
	db.DB.Find(&user, id)

	c.JSON(http.StatusOK, user)
}

// CreateUser is for creating an user
func CreateUser(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
