package model

import (
	file_utils "crowdfunding-api/src/utils/file-utils"
	"mime/multipart"
)

type FormUploadRequest struct {
	File      multipart.File
	FileInfo  file_utils.FileInfo
	Path      string `form:"path" binding:"required"`
	Overwrite bool   `form:"overwrite"`
	PrevFile  string `form:"previous_file"`
}

type UploadMultipleFileRequest struct {
	Files     []multipart.File
	FileInfo  []file_utils.FileInfo
	Path      string   `form:"path" binding:"required"`
	Overwrite bool     `form:"overwrite"`
	PrevFiles []string `form:"previous_files"`
}

type UploadResponse struct {
	FileUrl string `json:"file_url"`
}

type UploadMultipleFileResponse struct {
	FilesUrl []string `json:"files_url"`
}
