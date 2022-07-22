package controller

import (
	"crowdfunding-api/src/modules/user/model/dto"
	"crowdfunding-api/src/modules/user/service"
	"crowdfunding-api/src/utils"
	"crowdfunding-api/src/utils/common"
	e "crowdfunding-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (controller *UserControllerImpl) GetAll(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	data, err := controller.Service.GetAll(userClaims)
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get All User Success", data)
}

func (controller *UserControllerImpl) GetProfile(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	data, err := controller.Service.GetProfile(uint(userID))
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get User Profile Success", data)
}

func (controller *UserControllerImpl) GetUserCampaigns(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	data, err := controller.Service.GetUserCampaigns(uint(userID))
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get User Campaigns Success", data)
}

func (controller *UserControllerImpl) Update(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.UpdateUserRequest
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

	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	err = controller.Service.Update(userClaims, request, uint(userID))
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

	common.SendSuccess(ctx, http.StatusOK, "Update User Success", nil)
}

func (controller *UserControllerImpl) UpdateRole(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.UpdateUserRoleRequest
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

	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	err = controller.Service.UpdateRole(userClaims, request, uint(userID))
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

	common.SendSuccess(ctx, http.StatusOK, "Update User Role Success", nil)
}

func (controller *UserControllerImpl) Delete(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	err = controller.Service.Delete(userClaims, uint(userID))
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

	common.SendSuccess(ctx, http.StatusOK, "Delete User Success", nil)
}
