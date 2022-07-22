package repository

import (
	"crowdfunding-api/src/modules/campaign/model/entity"
	e "crowdfunding-api/src/utils/errors"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

type CampaignRepositoryImpl struct {
	DB *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &CampaignRepositoryImpl{DB: db}
}

func (repository *CampaignRepositoryImpl) Create(campaign entity.Campaign) error {
	result := repository.DB.Create(&campaign)

	var mySqlErr *mysql.MySQLError
	if errors.As(result.Error, &mySqlErr) && mySqlErr.Number == 1062 {
		result.Error = e.ErrDuplicateKey
	}

	return result.Error
}

func (repository *CampaignRepositoryImpl) Update(campaign entity.Campaign, campaignID uint) error {
	result := repository.DB.
		Model(&entity.Campaign{}).
		Where("id=?", campaignID).
		Updates(map[string]interface{}{
			"category_id":   campaign.CategoryID,
			"title":         campaign.Title,
			"deadline":      campaign.Deadline,
			"target_amount": campaign.TargetAmount,
			"image":         campaign.Image,
			"description":   campaign.Description,
		})
	return result.Error
}

func (repository *CampaignRepositoryImpl) UpdateCurrentAmount(campaign entity.Campaign, campaignID uint) error {
	result := repository.DB.
		Model(&entity.Campaign{}).
		Where("id=?", campaignID).
		Updates(map[string]interface{}{
			"current_amount": campaign.CurrentAmount,
			"is_completed":   campaign.IsCompleted,
		})
	return result.Error
}

func (repository *CampaignRepositoryImpl) Delete(campaignID uint, deleteBy string) error {
	result := repository.DB.
		Model(&entity.Campaign{}).
		Where("id=?", campaignID).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"deleted_by": deleteBy,
		})
	return result.Error
}

func (repository *CampaignRepositoryImpl) FindAll() ([]entity.CampaignLite, error) {
	var campaigns []entity.CampaignLite

	query := `
	SELECT c.id, c.title, c.image, u.name creator, cat.name category, c.deadline,
		c.target_amount, c.current_amount, c.is_completed
	FROM campaigns c
	LEFT JOIN users u ON u.id = c.user_id
	LEFT JOIN categories cat ON cat.id = c.category_id
	WHERE c.deleted_at IS NULL
	`
	result := repository.DB.Raw(query).Scan(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}

	return campaigns, nil
}

func (repository *CampaignRepositoryImpl) FindOne(campaignID uint) (entity.Campaign, error) {
	var campaign entity.Campaign
	result := repository.DB.Where("id=?", campaignID).First(&campaign)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrDataNotFound
		}
		return campaign, result.Error
	}
	return campaign, nil
}

func (repository *CampaignRepositoryImpl) FindOneDetail(campaignID uint) (entity.CampaignDetail, error) {
	var campaign entity.CampaignDetail

	query := `
	SELECT c.*, u.name creator, cat.name category
	FROM campaigns c
	LEFT JOIN users u ON u.id = c.user_id
	LEFT JOIN categories cat ON cat.id = c.category_id
	WHERE c.id = ?
	LIMIT 1
	`
	result := repository.DB.Raw(query, campaignID).Scan(&campaign)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrDataNotFound
		}
		return campaign, result.Error
	}

	return campaign, nil
}

func (repository *CampaignRepositoryImpl) FindByUserID(userID uint) ([]entity.CampaignLite, error) {
	var campaigns []entity.CampaignLite

	query := `
	SELECT c.id, c.title, c.image, u.name creator, cat.name category, c.deadline,
		c.target_amount, c.current_amount, c.is_completed
	FROM campaigns c
	LEFT JOIN users u ON u.id = c.user_id
	LEFT JOIN categories cat ON cat.id = c.category_id
	WHERE c.deleted_at IS NULL AND c.user_id = ?
	`
	result := repository.DB.Raw(query, userID).Scan(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}

	return campaigns, nil
}
