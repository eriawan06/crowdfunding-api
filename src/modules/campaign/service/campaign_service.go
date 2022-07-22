package service

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/campaign/model/dto"
)

type CampaignService interface {
	Create(claims ad.UserClaims, request dto.CreateCampaignRequest) error
	Update(claims ad.UserClaims, request dto.UpdateCampaignRequest, campaignID uint) error
	Delete(claims ad.UserClaims, campaignID uint) error
	GetAll() ([]dto.CampaignResponse, error)
	GetOne(campaignID uint) (dto.CampaignDetailResponse, error)
}
