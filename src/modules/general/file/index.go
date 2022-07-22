package file

import (
	"crowdfunding-api/src/modules/general/file/controller"
	"crowdfunding-api/src/modules/general/file/service"
)

var (
	fileController controller.FileController
	fileService    service.FileService
)

type Module interface {
	InitModule()
}

type ModuleImpl struct{}

func New() Module {
	return &ModuleImpl{}
}

func (module ModuleImpl) InitModule() {
	fileService = service.NewFileServiceCloudinaryImpl()
	fileController = controller.NewFileControllerImpl(fileService)
}

func GetFileController() controller.FileController {
	return fileController
}
