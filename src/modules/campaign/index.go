package campaign

import (
	"crowdfunding-api/src/modules/campaign/controller"
	"crowdfunding-api/src/modules/campaign/repository"
	"crowdfunding-api/src/modules/campaign/service"
	"crowdfunding-api/src/utils/common"
	"gorm.io/gorm"
)

var (
	campaignController controller.CampaignController
)

type ModuleImplCampaign struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplCampaign{DB: database}
}

func (module *ModuleImplCampaign) InitModule() {
	campaignRepository := repository.NewCampaignRepository(module.DB)
	campaignService := service.NewCampaignService(campaignRepository)
	campaignController = controller.NewCampaignController(campaignService)
}

func GetController() controller.CampaignController {
	return campaignController
}
