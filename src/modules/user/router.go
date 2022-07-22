package user

import "github.com/gin-gonic/gin"

func NewRouter(group *gin.RouterGroup) {
	controller := GetController()

	group.GET("/", controller.GetAll)
	group.GET("/:id", controller.GetProfile)
	group.GET("/:id/campaigns", controller.GetUserCampaigns)
	group.PUT("/:id", controller.Update)
	group.PUT("/:id/update-role", controller.UpdateRole)
	group.DELETE("/:id", controller.Delete)
}
