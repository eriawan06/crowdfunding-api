package service

import (
	"context"
	"crowdfunding-api/src/modules/general/file/model"
	"crowdfunding-api/src/utils"
	e "crowdfunding-api/src/utils/errors"
	file_utils "crowdfunding-api/src/utils/file-utils"
	"crowdfunding-api/src/utils/helper"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
	"os"
)

type FileServiceCloudinaryImpl struct{}

func NewFileServiceCloudinaryImpl() FileService {
	return &FileServiceCloudinaryImpl{}
}

func (service *FileServiceCloudinaryImpl) Upload(
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
			"%s%s",
			utils.GenerateRandomAlphaNumberic(10),
			request.FileInfo.FileExt)
	}

	secret := os.Getenv("CLOUDINARY_API_SECRET")
	cld, _ := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		secret,
	)

	result, err := cld.Upload.Upload(ctx, request.File, uploader.UploadParams{
		Folder:           request.Path,
		FilenameOverride: request.FileInfo.FileName,
		UseFilename:      helper.ReferBool(true),
		UniqueFilename:   helper.ReferBool(false),
		Overwrite:        helper.ReferBool(request.Overwrite),
	})

	if err != nil {
		return
	}

	response = model.UploadResponse{FileUrl: result.SecureURL}
	return
}

func (service *FileServiceCloudinaryImpl) UploadAsync(
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
			"%s%s",
			utils.GenerateRandomAlphaNumberic(10),
			request.FileInfo.FileExt)
	}

	uploadRespCh := make(chan model.UploadResponse, 10)
	errCh := make(chan error, 10)

	defer close(uploadRespCh)
	defer close(errCh)

	// Cloudinary credentials.
	secret := os.Getenv("CLOUDINARY_API_SECRET")
	cld, _ := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		secret,
	)

	go func(uploadResp chan model.UploadResponse, err chan error) {
		result, errUpload := cld.Upload.Upload(ctx, request.File, uploader.UploadParams{
			Folder:           request.Path,
			FilenameOverride: request.FileInfo.FileName,
			UseFilename:      helper.ReferBool(true),
			UniqueFilename:   helper.ReferBool(false),
			Overwrite:        helper.ReferBool(request.Overwrite),
		})

		uploadResp <- model.UploadResponse{FileUrl: result.SecureURL}
		err <- errUpload

	}(uploadRespCh, errCh)

	if err = <-errCh; err != nil {
		return
	}

	response = <-uploadRespCh
	return
}

func (service *FileServiceCloudinaryImpl) UploadMultipleFiles(
	ctx context.Context,
	request model.UploadMultipleFileRequest,
) (response model.UploadMultipleFileResponse, err error) {

	if request.Overwrite {
		if len(request.FileInfo) != len(request.PrevFiles) {
			err = errors.New("can't overwrite files, due to prev files length is not equal to files length")
			return
		}
	}

	errCh := make(chan error, 10)
	defer close(errCh)

	for k, info := range request.FileInfo {
		//check file format or extension
		if !helper.StringInSlice(info.FileExt, file_utils.ListAllowedFormatFile) {
			err = e.ErrUnsupportedFileFormat
			return
		}

		if !request.Overwrite {
			info.FileName = fmt.Sprintf(
				"%s%s",
				utils.GenerateRandomAlphaNumberic(10),
				info.FileExt,
			)
		} else {
			info.FileName = request.PrevFiles[k]
		}

		request.FileInfo[k] = info
	}

	uploadRespCh := make(chan model.UploadResponse, 10)
	defer close(uploadRespCh)

	// Cloudinary credentials.
	secret := os.Getenv("CLOUDINARY_API_SECRET")
	cld, _ := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		secret,
	)

	for k, file := range request.Files {
		go func(
			file multipart.File,
			fileInfo file_utils.FileInfo,
			overwrite bool,
			uploadResp chan model.UploadResponse,
			err chan error,
		) {
			result, errUpload := cld.Upload.Upload(ctx, file, uploader.UploadParams{
				Folder:           request.Path,
				FilenameOverride: fileInfo.FileName,
				UseFilename:      helper.ReferBool(true),
				UniqueFilename:   helper.ReferBool(false),
				Overwrite:        helper.ReferBool(overwrite),
			})

			uploadResp <- model.UploadResponse{FileUrl: result.SecureURL}
			err <- errUpload

		}(file, request.FileInfo[k], request.Overwrite, uploadRespCh, errCh)

		if err = <-errCh; err != nil {
			return
		}

		uploadResp := <-uploadRespCh
		response.FilesUrl = append(response.FilesUrl, uploadResp.FileUrl)
	}

	return
}
