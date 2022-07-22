package campaign

import "github.com/gin-gonic/gin"

func NewRouter(group *gin.RouterGroup) {
	group.POST("/", GetController().Create)
	group.PUT("/:id", GetController().Update)
	group.DELETE("/:id", GetController().Delete)
}

func NewPublicRouter(group *gin.RouterGroup) {
	group.GET("/", GetController().GetAll)
	group.GET("/:id", GetController().GetOne)
}
