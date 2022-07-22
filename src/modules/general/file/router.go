package file

import (
	"github.com/gin-gonic/gin"
)

func NewFileRouter(group *gin.RouterGroup) {
	group.POST("/upload", GetFileController().Upload)
	//group.POST("/upload-multiple", GetFileController().UploadMultipleFiles)
}
