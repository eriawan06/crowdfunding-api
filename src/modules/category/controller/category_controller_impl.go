package controller

import (
	"crowdfunding-api/src/modules/category/model/dto"
	"crowdfunding-api/src/modules/category/service"
	"crowdfunding-api/src/utils"
	"crowdfunding-api/src/utils/common"
	e "crowdfunding-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{Service: service}
}

func (controller *CategoryControllerImpl) Create(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.CreateCategoryRequest
	errorBinding := ctx.ShouldBindJSON(&request)
	if errorBinding != nil {
		// Check if there is EOF error
		if errorBinding.Error() == "EOF" {
			common.SendError(ctx, http.StatusBadRequest, "Body is empty", []string{"Body required"})
			return
		}

		// When Binding Error
		common.SendError(ctx, http.StatusBadRequest, "Invalid request", utils.SplitError(errorBinding))
		return
	}

	// Access Services
	err = controller.Service.Create(userClaims, request)
	if err != nil {
		if err == e.ErrDuplicateKey {
			common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{err.Error()})
			return
		}

		if err == e.ErrForbidden {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusCreated, "Create Category Success", nil)
}

func (controller *CategoryControllerImpl) Update(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.UpdateCategoryRequest
	errorBinding := ctx.ShouldBindJSON(&request)
	if errorBinding != nil {
		// Check if there is EOF error
		if errorBinding.Error() == "EOF" {
			common.SendError(ctx, http.StatusBadRequest, "Body is empty", []string{"Body required"})
			return
		}

		// When Binding Error
		common.SendError(ctx, http.StatusBadRequest, "Invalid request", utils.SplitError(errorBinding))
		return
	}

	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	// Access Services
	err = controller.Service.Update(userClaims, request, uint(categoryId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		if err == e.ErrForbidden {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Update Category Success", nil)
}

func (controller *CategoryControllerImpl) Delete(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	// Access Service
	err = controller.Service.Delete(userClaims, uint(categoryId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		if err == e.ErrForbidden {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Delete Category Success", nil)
}

func (controller *CategoryControllerImpl) GetAll(ctx *gin.Context) {
	var filter dto.FilterCategory
	filter.Name = ctx.Query("name")
	filter.IsActive = ctx.Query("is_active")

	data, err := controller.Service.GetAll(filter)
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get All Category Success", data)
}

func (controller *CategoryControllerImpl) GetOne(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	data, err := controller.Service.GetOne(uint(categoryId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get One Category Success", data)
}
