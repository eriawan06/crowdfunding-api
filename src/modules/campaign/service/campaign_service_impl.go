package service

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/campaign/mapper"
	"crowdfunding-api/src/modules/campaign/model/dto"
	"crowdfunding-api/src/modules/campaign/repository"
	e "crowdfunding-api/src/utils/errors"
	"crowdfunding-api/src/utils/helper"
	"time"
)

type CampaignServiceImpl struct {
	Repository repository.CampaignRepository
}

func NewCampaignService(repository repository.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{
		Repository: repository,
	}
}

func (service *CampaignServiceImpl) Create(claims ad.UserClaims, request dto.CreateCampaignRequest) error {
	if claims.Role != "campaigner" {
		return e.ErrForbidden
	}

	//create campaign
	campaign, err := mapper.CreateCampaignRequestToCampaign(claims, request)
	if err != nil {
		return err
	}

	err = service.Repository.Create(campaign)
	if err != nil {
		return err
	}

	return nil
}

func (service *CampaignServiceImpl) Update(claims ad.UserClaims, request dto.UpdateCampaignRequest, campaignID uint) error {
	if claims.Role != "campaigner" {
		return e.ErrForbidden
	}

	// check campaign
	checkCampaign, err := service.Repository.FindOne(campaignID)
	if err != nil {
		return err
	}

	// check campaign's owner
	if checkCampaign.UserID != claims.UserId {
		return e.ErrNotTheOwner
	}

	// update campaign
	campaign, err := mapper.UpdateCampaignRequestToCampaign(claims, request)
	if err != nil {
		return err
	}

	err = service.Repository.Update(campaign, campaignID)
	return err
}

func (service *CampaignServiceImpl) Delete(claims ad.UserClaims, campaignID uint) error {
	if !helper.StringInSlice(claims.Role, []string{"campaigner", "admin"}) {
		return e.ErrForbidden
	}

	// check campaign
	checkCampaign, err := service.Repository.FindOne(campaignID)
	if err != nil {
		return err
	}

	// check campaign's owner
	if claims.Role == "campaigner" && checkCampaign.UserID != claims.UserId {
		return e.ErrNotTheOwner
	}

	// delete campaign
	err = service.Repository.Delete(campaignID, claims.Email)
	return err
}

func (service *CampaignServiceImpl) GetAll() ([]dto.CampaignResponse, error) {
	campaigns, err := service.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var campaignsResponse []dto.CampaignResponse
	for _, campaign := range campaigns {
		timeRemain := campaign.Deadline.Sub(time.Now()).Hours()
		dayRemain := timeRemain / 24
		campaignsResponse = append(campaignsResponse, mapper.CampaignLiteToCampaignResponse(campaign, dayRemain))
	}
	return campaignsResponse, err
}

func (service *CampaignServiceImpl) GetOne(campaignID uint) (dto.CampaignDetailResponse, error) {
	campaign, err := service.Repository.FindOneDetail(campaignID)
	if err != nil {
		return dto.CampaignDetailResponse{}, err
	}

	timeRemain := campaign.Deadline.Sub(time.Now()).Hours()
	dayRemain := timeRemain / 24
	campaignResponse := mapper.CampaignDetailToCampaignDetailResponse(campaign, dayRemain)
	return campaignResponse, err
}
