package service

import (
	"crowdfunding-api/src/modules/auth/mapper"
	"crowdfunding-api/src/modules/auth/model/dto"
	ud "crowdfunding-api/src/modules/user/model/dto"
	ue "crowdfunding-api/src/modules/user/model/entity"
	ur "crowdfunding-api/src/modules/user/repository"
	"crowdfunding-api/src/utils"
	e "crowdfunding-api/src/utils/errors"
	"errors"
)

type AuthServiceImpl struct {
	UserRepo ur.UserRepository
}

func NewAuthService(userRepo ur.UserRepository) AuthService {
	return &AuthServiceImpl{UserRepo: userRepo}
}

func (service *AuthServiceImpl) Register(request dto.RegisterRequest) error {
	if request.Password != request.ConfirmPassword {
		return e.ErrConfirmPasswordNotSame
	}

	hashed, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}
	request.Password = hashed

	// mapping
	user := mapper.RegisterRequestToUser(request)

	// Create new User
	err = service.UserRepo.Create(user)
	return err
}

func (service *AuthServiceImpl) Login(request dto.LoginRequest) (dto.AuthResponse, error) {
	var (
		user ue.User
		err  error
	)

	user, err = service.UserRepo.FindByEmail(request.Email)
	if err != nil {
		if err.Error() == "record not found" {
			return dto.AuthResponse{}, e.ErrWrongLoginCredential
		}
		return dto.AuthResponse{}, err
	}

	isPasswordValid := utils.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordValid {
		return dto.AuthResponse{}, e.ErrWrongLoginCredential
	}

	// Generate Token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// Return nil Error
	return dto.AuthResponse{
		Token: token,
		User: ud.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}

func (service *AuthServiceImpl) GoogleOauth(request utils.GoogleUserResult) (dto.AuthResponse, error) {
	var (
		user ue.User
		err  error
	)

	user, err = service.UserRepo.FindByEmail(request.Email)
	if err != nil && err.Error() != "email is not registered" {
		return dto.AuthResponse{}, err
	}

	// if user is not exist (user.id = 0), then create user with auth type google
	if user.Id == 0 {
		err = service.UserRepo.Create(mapper.GoogleUserResultToUser(request))
		if err != nil {
			return dto.AuthResponse{}, err
		}

		user, err = service.UserRepo.FindByEmail(request.Email)
		if err != nil && err.Error() != "email is not registered" {
			return dto.AuthResponse{}, err
		}
	}

	// check auth type of existing user
	if user.AuthType != "google" {
		return dto.AuthResponse{}, errors.New("email was registered by other auth method")
	}

	// Generate Token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// Return nil Error
	return dto.AuthResponse{
		Token: token,
		User: ud.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
