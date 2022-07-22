package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAll(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	GetUserCampaigns(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
