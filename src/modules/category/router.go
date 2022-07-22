package category

import "github.com/gin-gonic/gin"

func NewRouter(group *gin.RouterGroup) {
	controller := GetController()
	group.POST("", controller.Create)
	group.PUT("/:id", GetController().Update)
	group.DELETE("/:id", GetController().Delete)
}

func NewPublicRouter(group *gin.RouterGroup) {
	group.GET("", GetController().GetAll)
	group.GET("/:id", GetController().GetOne)
}
