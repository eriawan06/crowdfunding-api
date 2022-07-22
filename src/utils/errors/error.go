package errors

import "errors"

var (
	// Custom Error
	ErrDuplicateKey           = errors.New("duplicate key")
	ErrDataNotFound           = errors.New("data not found")
	ErrForbidden              = errors.New("you don't have permission to access")
	ErrNotTheOwner            = errors.New("you are not the owner of this resource")
	ErrWrongLoginCredential   = errors.New("wrong email or password")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrEmailNotRegistered     = errors.New("email is not registered")
	ErrConfirmPasswordNotSame = errors.New("password and confirm password are not same")
	ErrUnsupportedFileFormat  = errors.New("file format is not supported")
	ErrWrongFileUploadPath    = errors.New("wrong file upload path")
)
