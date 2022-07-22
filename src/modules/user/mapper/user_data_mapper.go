package mapper

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/user/model/dto"
	"crowdfunding-api/src/modules/user/model/entity"
	"crowdfunding-api/src/utils/common"
	"crowdfunding-api/src/utils/helper"
)

func UpdateUserRequestToUser(claims ad.UserClaims, request dto.UpdateUserRequest) (entity.User, error) {
	parsedBirtDate, err := helper.ParseDateStringToTime(request.BirthDate)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		Name:      request.Name,
		BirthDate: &parsedBirtDate,
		Avatar:    request.Avatar,
		Address:   request.Address,
		Bio:       request.Bio,
		AuthType:  "regular",
		BaseEntity: common.BaseEntity{
			UpdatedBy: claims.Email,
		},
	}, nil
}
