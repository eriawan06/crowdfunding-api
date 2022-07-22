package service

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	campMapper "crowdfunding-api/src/modules/campaign/mapper"
	"crowdfunding-api/src/modules/campaign/model/dto"
	campRepo "crowdfunding-api/src/modules/campaign/repository"
	"crowdfunding-api/src/modules/user/mapper"
	ud "crowdfunding-api/src/modules/user/model/dto"
	"crowdfunding-api/src/modules/user/model/entity"
	"crowdfunding-api/src/modules/user/repository"
	"crowdfunding-api/src/utils/common"
	e "crowdfunding-api/src/utils/errors"
	"time"
)

type UserServiceImpl struct {
	Repository   repository.UserRepository
	CampaignRepo campRepo.CampaignRepository
}

func NewUserService(
	repo repository.UserRepository,
	campaignRepo campRepo.CampaignRepository,
) UserService {
	return &UserServiceImpl{
		Repository:   repo,
		CampaignRepo: campaignRepo,
	}
}

func (service *UserServiceImpl) GetAll(claims ad.UserClaims) (users []ud.UserResponse, err error) {
	if claims.Role != "admin" {
		err = e.ErrForbidden
		return
	}

	usersDb, err := service.Repository.FindAll()
	if err != nil {
		return
	}

	for _, user := range usersDb {
		users = append(users, ud.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}
	return
}

func (service *UserServiceImpl) GetProfile(userID uint) (user ud.UserProfileResponse, err error) {
	userDb, err := service.Repository.FindById(userID)
	if err != nil {
		return
	}

	user = ud.UserProfileResponse{
		ID:        userDb.Id,
		Name:      userDb.Name,
		Email:     userDb.Email,
		Role:      userDb.Role,
		BirthDate: userDb.BirthDate,
		Avatar:    userDb.Avatar,
		Address:   userDb.Address,
		Bio:       userDb.Bio,
	}
	return
}

func (service *UserServiceImpl) GetUserCampaigns(userID uint) (campaigns []dto.CampaignResponse, err error) {
	campaignsDb, err := service.CampaignRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	for _, campaign := range campaignsDb {
		timeRemain := campaign.Deadline.Sub(time.Now()).Hours()
		dayRemain := timeRemain / 24
		campaigns = append(campaigns, campMapper.CampaignLiteToCampaignResponse(campaign, dayRemain))
	}
	return campaigns, err
}

func (service *UserServiceImpl) Update(claims ad.UserClaims, request ud.UpdateUserRequest, userID uint) (err error) {
	//check
	if _, err = service.Repository.FindById(userID); err != nil {
		return
	}

	//mapping
	user, err := mapper.UpdateUserRequestToUser(claims, request)
	if err != nil {
		return err
	}

	if err = service.Repository.Update(user, userID); err != nil {
		return
	}

	return
}

func (service *UserServiceImpl) UpdateRole(claims ad.UserClaims, request ud.UpdateUserRoleRequest, userID uint) (err error) {
	//check
	if _, err = service.Repository.FindById(userID); err != nil {
		return
	}

	user := entity.User{
		Role: request.Role,
		BaseEntity: common.BaseEntity{
			UpdatedBy: claims.Email,
		}}
	if err = service.Repository.UpdateRole(user, userID); err != nil {
		return
	}

	return
}

func (service *UserServiceImpl) Delete(claims ad.UserClaims, userID uint) (err error) {
	if claims.Role != "admin" {
		err = e.ErrForbidden
		return
	}

	//check
	if _, err = service.Repository.FindById(userID); err != nil {
		return
	}

	if err = service.Repository.Delete(userID, claims.Email); err != nil {
		return
	}

	return
}
