package repository

import "crowdfunding-api/src/modules/campaign/model/entity"

type CampaignRepository interface {
	Create(campaign entity.Campaign) error
	Update(campaign entity.Campaign, campaignID uint) error
	UpdateCurrentAmount(campaign entity.Campaign, campaignID uint) error
	Delete(campaignID uint, deleteBy string) error
	FindAll() ([]entity.CampaignLite, error)
	FindOne(campaignID uint) (entity.Campaign, error)
	FindOneDetail(campaignID uint) (entity.CampaignDetail, error)
	FindByUserID(userID uint) ([]entity.CampaignLite, error)
}
