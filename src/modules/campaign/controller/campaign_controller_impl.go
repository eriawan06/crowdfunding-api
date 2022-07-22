package controller

import (
	"crowdfunding-api/src/modules/campaign/model/dto"
	"crowdfunding-api/src/modules/campaign/service"
	"crowdfunding-api/src/utils"
	"crowdfunding-api/src/utils/common"
	e "crowdfunding-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CampaignControllerImpl struct {
	Service service.CampaignService
}

func NewCampaignController(service service.CampaignService) CampaignController {
	return &CampaignControllerImpl{Service: service}
}

func (controller CampaignControllerImpl) Create(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.CreateCampaignRequest
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

	err = controller.Service.Create(userClaims, request)
	if err != nil {
		if err == e.ErrForbidden {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusCreated, "Create Campaign Success", nil)
}

func (controller CampaignControllerImpl) Update(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.UpdateCampaignRequest
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

	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	err = controller.Service.Update(userClaims, request, uint(courseId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		if err == e.ErrForbidden || err == e.ErrNotTheOwner {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Update Campaign Success", nil)
}

func (controller CampaignControllerImpl) Delete(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	err = controller.Service.Delete(userClaims, uint(courseId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		if err == e.ErrForbidden || err == e.ErrNotTheOwner {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Delete Campaign Success", nil)
}

func (controller CampaignControllerImpl) GetAll(ctx *gin.Context) {
	data, err := controller.Service.GetAll()
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get All Campaign Success", data)
}

func (controller CampaignControllerImpl) GetOne(ctx *gin.Context) {
	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	data, err := controller.Service.GetOne(uint(courseId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get One Campaign Success", data)
}
