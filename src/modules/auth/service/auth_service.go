package service

import (
	"crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/utils"
)

type AuthService interface {
	Register(request dto.RegisterRequest) error
	Login(request dto.LoginRequest) (dto.AuthResponse, error)
	GoogleOauth(request utils.GoogleUserResult) (dto.AuthResponse, error)
}
