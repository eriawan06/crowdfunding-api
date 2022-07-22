package service

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	cd "crowdfunding-api/src/modules/campaign/model/dto"
	"crowdfunding-api/src/modules/user/model/dto"
)

type UserService interface {
	GetAll(claims ad.UserClaims) (users []dto.UserResponse, err error)
	GetProfile(userID uint) (user dto.UserProfileResponse, err error)
	GetUserCampaigns(userID uint) (campaigns []cd.CampaignResponse, err error)
	Update(
		claims ad.UserClaims,
		request dto.UpdateUserRequest,
		userID uint,
	) (err error)
	UpdateRole(
		claims ad.UserClaims,
		request dto.UpdateUserRoleRequest,
		userID uint,
	) (err error)
	Delete(claims ad.UserClaims, userID uint) (err error)
}
