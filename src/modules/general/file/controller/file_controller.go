package controller

import "github.com/gin-gonic/gin"

type FileController interface {
	Upload(ctx *gin.Context)
	UploadMultipleFiles(ctx *gin.Context)
}
