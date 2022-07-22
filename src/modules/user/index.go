package user

import (
	campRepo "crowdfunding-api/src/modules/campaign/repository"
	"crowdfunding-api/src/modules/user/controller"
	"crowdfunding-api/src/modules/user/repository"
	"crowdfunding-api/src/modules/user/service"
	"crowdfunding-api/src/utils/common"
	"gorm.io/gorm"
)

var (
	userRepository repository.UserRepository
	userController controller.UserController
)

type ModuleImplAuth struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplAuth{DB: database}
}

func (module *ModuleImplAuth) InitModule() {
	userRepository = repository.NewUserRepository(module.DB)
	campaignRepo := campRepo.NewCampaignRepository(module.DB)
	userService := service.NewUserService(userRepository, campaignRepo)
	userController = controller.NewUserController(userService)
}

func GetRepository() repository.UserRepository {
	return userRepository
}

func GetController() controller.UserController {
	return userController
}
