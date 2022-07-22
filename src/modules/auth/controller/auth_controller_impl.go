package controller

import (
	"crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/auth/service"
	"crowdfunding-api/src/utils"
	"crowdfunding-api/src/utils/common"
	e "crowdfunding-api/src/utils/errors"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthControllerImpl struct {
	Service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &AuthControllerImpl{Service: service}
}

func (controller *AuthControllerImpl) Register(ctx *gin.Context) {
	var request dto.RegisterRequest

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

	//access register service
	err := controller.Service.Register(request)
	if err != nil {
		if err == e.ErrEmailAlreadyExists || err == e.ErrConfirmPasswordNotSame {
			common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Register Success", nil)
}

func (controller *AuthControllerImpl) Login(ctx *gin.Context) {
	var request dto.LoginRequest

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
	response, err := controller.Service.Login(request)
	if err != nil {
		if err == e.ErrWrongLoginCredential {
			common.SendError(ctx, http.StatusUnauthorized, "Unauthorized", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Login Success", &response)
}

func (controller *AuthControllerImpl) GoogleOauth(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.SetCookie("logged_in", "false", 1000*60, "/oauth/google/callback", "localhost", false, false)
		ctx.SetCookie("error", "Authorization code not provided!", 1000*60, "/oauth/google/callback", "localhost", true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/oauth/google/callback")
		//common.SendError(ctx, http.StatusUnauthorized, "Unauthorized", []string{"Authorization code not provided!"})
		return
	}

	// Use the code to get the id and access tokens
	tokenRes, err := utils.GetGoogleOauthToken(code)
	if err != nil {
		ctx.SetCookie("logged_in", "false", 1000*60, "/oauth/google/callback", "localhost", false, false)
		ctx.SetCookie("error", err.Error(), 1000*60, "/oauth/google/callback", "localhost", true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/oauth/google/callback")
		//common.SendError(ctx, http.StatusBadGateway, "Bad Gateway", []string{err.Error()})
		return
	}

	googleUser, err := utils.GetGoogleUser(tokenRes.AccessToken, tokenRes.IdToken)
	if err != nil {
		ctx.SetCookie("logged_in", "false", 1000*60, "/oauth/google/callback", "localhost", false, false)
		ctx.SetCookie("error", err.Error(), 1000*60, "/oauth/google/callback", "localhost", true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/oauth/google/callback")
		//common.SendError(ctx, http.StatusBadGateway, "Bad Gateway", []string{err.Error()})
		return
	}

	response, err := controller.Service.GoogleOauth(*googleUser)
	if err != nil {
		ctx.SetCookie("logged_in", "false", 1000*60, "/oauth/google/callback", "localhost", false, false)
		ctx.SetCookie("error", err.Error(), 1000*60, "/oauth/google/callback", "localhost", true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/oauth/google/callback")
		return
	}

	userData, err := json.Marshal(response.User)
	if err != nil {
		ctx.SetCookie("logged_in", "false", 1000*60, "/oauth/google/callback", "localhost", false, false)
		ctx.SetCookie("error", err.Error(), 1000*60, "/oauth/google/callback", "localhost", true, true)
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/oauth/google/callback")
		return
	}

	ctx.SetCookie("access_token", response.Token, 1000*60, "/oauth/google/callback", "localhost", true, true)
	ctx.SetCookie("user", string(userData), 1000*60, "/oauth/google/callback", "localhost", true, true)
	ctx.SetCookie("logged_in", "true", 1000*60, "/oauth/google/callback", "localhost", false, false)
	ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/oauth/google/callback")
}
