package common

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	Id        uint       `gorm:"primaryKey"`
	CreatedAt time.Time  `gorm:"not null;autoCreateTime"`
	CreatedBy string     `gorm:"type:varchar(36);null;default:NULL"`
	UpdatedAt time.Time  `gorm:"not null;autoUpdateTime"`
	UpdatedBy string     `gorm:"type:varchar(36);null;default:NULL"`
	DeletedAt *time.Time `gorm:"default:NULL"`
	DeletedBy *string    `gorm:"type:varchar(36);null;default:NULL"`
}

type BaseDtoResponse struct {
	Id        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	DeletedBy sql.NullString `json:"deleted_by"`
}

type Module interface {
	InitModule()
}
