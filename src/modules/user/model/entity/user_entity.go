package entity

import (
	"crowdfunding-api/src/utils/common"
	"time"
)

type User struct {
	common.BaseEntity
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);null"`
	Role      string `gorm:"type:varchar(20);not null"`
	BirthDate *time.Time
	Avatar    *string `gorm:"type:varchar(255)"`
	Address   *string `gorm:"type:text"`
	Bio       *string `gorm:"type:text"`
	AuthType  string  `gorm:"type:varchar(20);not null;default:regular"`
}
