package auth

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(group *gin.RouterGroup) {
	group.POST("/register", authController.Register)
	group.POST("/login", authController.Login)
	//group.GET("/google", authController.GoogleOauth)
}
