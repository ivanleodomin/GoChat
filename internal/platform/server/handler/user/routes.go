package user_handler

import (
	"app-go/internal/platform/storage/postgresql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRepository := postgresql.NewUserRepository()

	userGroup := router.Group("/user")
	{
		userGroup.GET("/register", Register(userRepository))
	}
}
