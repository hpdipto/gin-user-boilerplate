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
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser is for creating an user
func CreateUser(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser is for updating an user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// retrieving user
	var user User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// taking input
	var inputUser User
	err = c.ShouldBindJSON(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update in DB
	db.DB.Model(&user).Updates(inputUser)

	c.JSON(http.StatusOK, user)
}

// DeleteUser for deleting a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	db.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
