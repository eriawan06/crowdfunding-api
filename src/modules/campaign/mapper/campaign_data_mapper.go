package mapper

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/campaign/model/dto"
	"crowdfunding-api/src/modules/campaign/model/entity"
	"crowdfunding-api/src/utils/common"
	"crowdfunding-api/src/utils/helper"
	"math"
)

func CreateCampaignRequestToCampaign(claims ad.UserClaims, request dto.CreateCampaignRequest) (entity.Campaign, error) {
	deadlineDate, err := helper.ParseDateStringToTime(request.Deadline)
	if err != nil {
		return entity.Campaign{}, err
	}

	return entity.Campaign{
		UserID:       claims.UserId,
		CategoryID:   request.CategoryID,
		Title:        request.Title,
		Deadline:     deadlineDate,
		TargetAmount: request.TargetAmount,
		Image:        request.Image,
		Description:  request.Description,
		BaseEntity: common.BaseEntity{
			CreatedBy: claims.Email,
			UpdatedBy: claims.Email,
		},
	}, nil
}

func UpdateCampaignRequestToCampaign(claims ad.UserClaims, request dto.UpdateCampaignRequest) (entity.Campaign, error) {
	deadlineDate, err := helper.ParseDateStringToTime(request.Deadline)
	if err != nil {
		return entity.Campaign{}, err
	}

	return entity.Campaign{
		CategoryID:   request.CategoryID,
		Title:        request.Title,
		Deadline:     deadlineDate,
		TargetAmount: request.TargetAmount,
		Image:        request.Image,
		Description:  request.Description,
		BaseEntity: common.BaseEntity{
			UpdatedBy: claims.Email,
		},
	}, nil
}

func CampaignLiteToCampaignResponse(cl entity.CampaignLite, dayRemain float64) dto.CampaignResponse {
	return dto.CampaignResponse{
		ID:            cl.ID,
		Title:         cl.Title,
		Image:         cl.Image,
		Creator:       cl.Creator,
		Category:      cl.Category,
		DayRemaining:  int(math.Round(dayRemain)),
		TargetAmount:  cl.TargetAmount,
		CurrentAmount: cl.CurrentAmount,
		IsCompleted:   *cl.IsCompleted,
	}
}

func CampaignDetailToCampaignDetailResponse(cd entity.CampaignDetail, dayRemain float64) dto.CampaignDetailResponse {
	return dto.CampaignDetailResponse{
		ID:            cd.ID,
		Title:         cd.Title,
		Image:         cd.Image,
		Creator:       cd.Creator,
		Category:      cd.Category,
		DayRemaining:  int(math.Round(dayRemain)),
		TargetAmount:  cd.TargetAmount,
		CurrentAmount: cd.CurrentAmount,
		Description:   cd.Description,
		IsCompleted:   *cd.IsCompleted,
		TotalDonation: cd.TotalDonation,
	}
}
