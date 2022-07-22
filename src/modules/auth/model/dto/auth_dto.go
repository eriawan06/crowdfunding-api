package dto

import (
	ud "crowdfunding-api/src/modules/user/model/dto"
)

type RegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
	Authorized bool   `json:"authorized"`
	UserId     uint   `json:"user_id"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Expired    int64  `json:"expired"`
}

type AuthResponse struct {
	Token string          `json:"token"`
	User  ud.UserResponse `json:"user"`
}
