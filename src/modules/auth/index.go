package auth

import (
	"crowdfunding-api/src/modules/auth/controller"
	"crowdfunding-api/src/modules/auth/service"
	ur "crowdfunding-api/src/modules/user/repository"
	"crowdfunding-api/src/utils/common"

	"gorm.io/gorm"
)

var (
	authService    service.AuthService
	authController controller.AuthController
)

type ModuleImplAuth struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplAuth{DB: database}
}

func (module *ModuleImplAuth) InitModule() {
	userRepository := ur.NewUserRepository(module.DB)
	authService = service.NewAuthService(userRepository)
	authController = controller.NewAuthController(authService)
}

func GetController() controller.AuthController {
	return authController
}
