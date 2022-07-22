package service

import (
	"context"
	"crowdfunding-api/src/modules/general/file/model"
	"crowdfunding-api/src/utils"
	e "crowdfunding-api/src/utils/errors"
	file_utils "crowdfunding-api/src/utils/file-utils"
	"crowdfunding-api/src/utils/helper"
	"fmt"
	"os"
)

type FileServiceS3Impl struct{}

func NewFileServiceS3Impl() FileService {
	return &FileServiceS3Impl{}
}

func (service *FileServiceS3Impl) Upload(
	ctx context.Context,
	request model.FormUploadRequest,
) (response model.UploadResponse, err error) {
	//check file format or extension
	if !helper.StringInSlice(request.FileInfo.FileExt, file_utils.ListAllowedFormatFile) {
		err = e.ErrUnsupportedFileFormat
		return
	}

	request.FileInfo.FileName = request.PrevFile
	if !request.Overwrite {
		request.FileInfo.FileName = fmt.Sprintf(
			"%s/%s%s",
			request.Path,
			utils.GenerateRandomAlphaNumberic(10),
			request.FileInfo.FileExt)
	}

	//// upload to s3
	if err = file_utils.PushS3(file_utils.S3Info{
		Endpoint: helper.ReferString(os.Getenv("AWS_S3_ENDPOINT")),
		Key:      os.Getenv("AWS_S3_ACCESS_KEY"),
		Secret:   os.Getenv("AWS_S3_SECRET_KEY"),
		Region:   os.Getenv("AWS_S3_REGION"),
		Bucket:   os.Getenv("AWS_S3_BUCKET"),
		File:     request.File,
		Filename: request.FileInfo.FileName,
		Filemime: request.FileInfo.FileMime,
		Filesize: request.FileInfo.FileSize,
	}); err != nil {
		return
	}

	response = model.UploadResponse{
		FileUrl: fmt.Sprintf("%s/%s", os.Getenv("AWS_S3_BASE_FILE_URL"), request.FileInfo.FileName),
	}
	return
}

func (service *FileServiceS3Impl) UploadAsync(ctx context.Context, request model.FormUploadRequest) (response model.UploadResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (service *FileServiceS3Impl) UploadMultipleFiles(ctx context.Context, request model.UploadMultipleFileRequest) (response model.UploadMultipleFileResponse, err error) {
	//TODO implement me
	panic("implement me")
}
