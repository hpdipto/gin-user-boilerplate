package user

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// GetAuthenticatedUser will return authenticated user's id
func GetAuthenticatedUser(c *gin.Context) (string, error) {

	// restoring token from session
	session := sessions.Default(c)
	sessionToken, ok := session.Get("token").(string)
	if !ok {
		return "", errors.New("authentication requied")
	}

	// parsing with claim, ignoring the token, taking id from the claim
	claims := make(jwt.MapClaims)
	_, jwterr := jwt.ParseWithClaims(sessionToken, claims, func(token *jwt.Token) (interface{}, error) {
		// ****** need to add this to env variable *******
		return []byte("secret"), nil
	})

	if jwterr != nil {
		return "", jwterr
	}

	return fmt.Sprint(claims["id"]), nil
}

// HashPassword for hasing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// ComparePasswordHash for comparing password and hash
func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
