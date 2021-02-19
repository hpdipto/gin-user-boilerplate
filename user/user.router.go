package user

import (
	"github.com/gin-gonic/gin"
)

// Routes for combining all user routes
func Routes(router *gin.Engine, baseURL string) {
	userRoute := router.Group(baseURL + "/user")
	{
		userRoute.GET("/:id", GetUser)
		userRoute.POST("/create", CreateUser)
	}
}
