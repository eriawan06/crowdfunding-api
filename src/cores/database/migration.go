package database

import (
	camEn "crowdfunding-api/src/modules/campaign/model/entity"
	catEn "crowdfunding-api/src/modules/category/model/entity"
	usEn "crowdfunding-api/src/modules/user/model/entity"
	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {
	db.AutoMigrate(&usEn.User{})
	db.AutoMigrate(&catEn.Category{})
	db.AutoMigrate(&camEn.Campaign{})
}
