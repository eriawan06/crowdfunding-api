package category

import (
	"crowdfunding-api/src/modules/category/controller"
	"crowdfunding-api/src/modules/category/repository"
	"crowdfunding-api/src/modules/category/service"
	"crowdfunding-api/src/utils/common"
	"gorm.io/gorm"
)

var (
	categoryController controller.CategoryController
)

type ModuleImplCategory struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplCategory{DB: database}
}

func (module *ModuleImplCategory) InitModule() {
	categoryRepository := repository.NewCategoryRepository(module.DB)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController = controller.NewCategoryController(categoryService)
}

func GetController() controller.CategoryController {
	return categoryController
}
