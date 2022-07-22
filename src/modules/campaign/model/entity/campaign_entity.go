package entity

import (
	catEn "crowdfunding-api/src/modules/category/model/entity"
	usEn "crowdfunding-api/src/modules/user/model/entity"
	"crowdfunding-api/src/utils/common"
	"time"
)

type Campaign struct {
	common.BaseEntity
	UserID        uint
	User          usEn.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoryID    uint
	Category      catEn.Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Title         string         `gorm:"type:varchar(60);not null"`
	Deadline      time.Time      `gorm:"not null"`
	TargetAmount  uint           `gorm:"not null"`
	CurrentAmount uint           `gorm:"not null;default: 0"`
	IsCompleted   *bool          `gorm:"not null;default: false"`
	Image         *string        `gorm:"type:varchar(255);null"`
	Description   *string        `gorm:"type:text"`
}

type CampaignLite struct {
	ID            uint
	Title         string
	Image         *string
	Creator       string
	Category      string
	Deadline      time.Time
	TargetAmount  uint
	CurrentAmount uint
	IsCompleted   *bool
}

type CampaignDetail struct {
	ID            uint
	Title         string
	Image         *string
	Creator       string
	Category      string
	Deadline      time.Time
	TargetAmount  uint
	CurrentAmount uint
	Description   *string
	IsCompleted   *bool
	TotalDonation uint
}
