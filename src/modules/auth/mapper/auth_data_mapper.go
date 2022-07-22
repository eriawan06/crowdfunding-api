package mapper

import (
	"crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/user/model/entity"
	"crowdfunding-api/src/utils"
	"crowdfunding-api/src/utils/common"
	"crowdfunding-api/src/utils/constants"
)

func RegisterRequestToUser(request dto.RegisterRequest) entity.User {
	return entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
		BaseEntity: common.BaseEntity{
			CreatedBy: "self",
			UpdatedBy: "self",
		},
	}
}

func GoogleUserResultToUser(request utils.GoogleUserResult) entity.User {
	return entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Role:     "user",
		AuthType: constants.AuthTypeGoogle,
		BaseEntity: common.BaseEntity{
			CreatedBy: "self",
			UpdatedBy: "self",
		},
	}
}
