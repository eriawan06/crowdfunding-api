package dto

import "time"

type UpdateUserRequest struct {
	Name      string  `json:"name"`
	BirthDate string  `json:"birth_date"`
	Avatar    *string `json:"avatar"`
	Address   *string `json:"address"`
	Bio       *string `json:"bio"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role"`
}

type UserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserProfileResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	BirthDate *time.Time `json:"birth_date"`
	Avatar    *string    `json:"avatar"`
	Address   *string    `json:"address"`
	Bio       *string    `json:"bio"`
}
