package controller

import (
	"crowdfunding-api/src/modules/general/file/model"
	"crowdfunding-api/src/modules/general/file/service"
	"crowdfunding-api/src/utils/common"
	e "crowdfunding-api/src/utils/errors"
	file_utils "crowdfunding-api/src/utils/file-utils"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type FileControllerImpl struct {
	Service service.FileService
}

func NewFileControllerImpl(fileService service.FileService) FileController {
	return &FileControllerImpl{
		Service: fileService,
	}
}

func (controller *FileControllerImpl) Upload(ctx *gin.Context) {
	// limit upload file size
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 10*file_utils.MB) // 1 Mb

	var request model.FormUploadRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{err.Error()})
		return
	}

	var multipartFileHeader *multipart.FileHeader
	request.File, multipartFileHeader, err = ctx.Request.FormFile("file")
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	// Create a buffer to store the header of the file in
	// And copy the headers into the FileHeader buffer
	fileHeader := make([]byte, 512)
	if _, err = request.File.Read(fileHeader); err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	// set position back to start.
	if _, err = request.File.Seek(0, 0); err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	defer request.File.Close()

	mType := mimetype.Detect(fileHeader)
	request.FileInfo = file_utils.FileInfo{
		FileName: multipartFileHeader.Filename,
		FileSize: request.File.(file_utils.Sizer).Size(),
		FileMime: mType.String(),
		FileExt:  mType.Extension(),
	}

	data, err := controller.Service.UploadAsync(ctx, request)
	if err != nil {
		if err == e.ErrUnsupportedFileFormat ||
			err == e.ErrWrongFileUploadPath {
			common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Upload File Success", data)
}

func (controller *FileControllerImpl) UploadMultipleFiles(ctx *gin.Context) {
	// limit upload file size
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 20*file_utils.MB) // 1 Mb

	var request model.UploadMultipleFileRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{err.Error()})
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	fileHeaders := form.File["files"]
	if fileHeaders == nil {
		common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{"Files should not be empty"})
		return
	}

	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
			return
		}

		//Create a buffer to store the header of the file in
		//And copy the headers into the FileHeader buffer
		fileHeaderBuffer := make([]byte, 512)
		if _, err = file.Read(fileHeaderBuffer); err != nil {
			common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
			return
		}

		// set position back to start.
		if _, err = file.Seek(0, 0); err != nil {
			common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
			return
		}

		defer file.Close()

		mType := mimetype.Detect(fileHeaderBuffer)
		fileInfo := file_utils.FileInfo{
			FileName: fileHeader.Filename,
			FileSize: file.(file_utils.Sizer).Size(),
			FileMime: mType.String(),
			FileExt:  mType.Extension(),
		}

		request.Files = append(request.Files, file)
		request.FileInfo = append(request.FileInfo, fileInfo)
	}

	data, err := controller.Service.UploadMultipleFiles(ctx, request)
	if err != nil {
		if err == e.ErrUnsupportedFileFormat ||
			err == e.ErrWrongFileUploadPath {
			common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Upload Multiple Files Success", data)
}
