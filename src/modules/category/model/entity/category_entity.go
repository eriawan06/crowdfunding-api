package entity

import "crowdfunding-api/src/utils/common"

type Category struct {
	common.BaseEntity
	Name     string `gorm:"type:varchar(255);uniqueIndex;not null"`
	IsActive *bool  `gorm:"default: true"`
}
