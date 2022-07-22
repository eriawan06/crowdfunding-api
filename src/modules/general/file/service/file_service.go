package service

import (
	"context"
	"crowdfunding-api/src/modules/general/file/model"
)

type FileService interface {
	Upload(
		ctx context.Context,
		request model.FormUploadRequest,
	) (response model.UploadResponse, err error)
	UploadAsync(
		ctx context.Context,
		request model.FormUploadRequest,
	) (response model.UploadResponse, err error)
	UploadMultipleFiles(
		ctx context.Context,
		request model.UploadMultipleFileRequest,
	) (response model.UploadMultipleFileResponse, err error)
}
