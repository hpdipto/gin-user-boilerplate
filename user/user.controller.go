package user

import (
	"fmt"
	db "gub/database"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
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
	var userInfo Info // this variable is for displaying non sensitive data

	// Get authenticated userID from session token
	userID, err := GetAuthenticatedUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if token id not matched with route id, reject the request
	if userID != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized user"})
		return
	}

	err = db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// displaying non sensitive data
	userInfo = Info(user)
	c.JSON(http.StatusCreated, userInfo)
}

// CreateUser is for creating an user
func CreateUser(c *gin.Context) {
	var user User
	var userInfo Info // this variable is for displaying non sensitive data

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hasing the password
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&user)
	// displaying non sensitive data
	userInfo = Info(user)
	c.JSON(http.StatusCreated, userInfo)
}

// UpdateUser is for updating an user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user User
	var userInfo Info // this variable is for displaying non sensitive data

	userID, err := GetAuthenticatedUser(c)
	// if token id not matched with route id, reject the request
	if userID != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized user"})
		return
	}

	// retrieving user
	err = db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// taking input
	var updatedUser User
	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update in DB without email and password field
	db.DB.Model(&user).Omit("email").Omit("password").Updates(updatedUser)
	// displaying non sensitive data
	userInfo = Info(user)
	c.JSON(http.StatusCreated, userInfo)
}

// Login functionality for a user
func Login(c *gin.Context) {
	var loginUser LoginUser
	var user User
	// var userInfo Info // this variable is for displaying non sensitive data

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find user by email
	db.DB.Where("email = ?", loginUser.Email).First(&user)
	if user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	// matching password
	// if matched then send the token to the user
	if ComparePasswordHash(loginUser.Password, user.Password) {
		// userInfo = Info(user)
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["id"] = user.ID
		// ekpires in 5 mins
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
		token.Claims = claims
		// ****** need to add this to env variable *******
		tokenString, jwterr := token.SignedString([]byte("secret"))
		if jwterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": jwterr.Error()})
			return
		}
		// saving the session
		session := sessions.Default(c)
		session.Set("token", tokenString)
		session.Save()
		c.JSON(http.StatusAccepted, gin.H{"authenticated": true, "token": tokenString})
		return

	}
	// if password don't match
	c.JSON(http.StatusBadRequest, gin.H{"error": "password doesn't match"})
}

// Logout for logging out a user
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

// DeleteUser for deleting a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	userID, err := GetAuthenticatedUser(c)
	// if token id not matched with route id, reject the request
	if userID != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized user"})
		return
	}

	var user User
	err = db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	db.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
