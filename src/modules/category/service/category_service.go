package service

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/category/model/dto"
)

type CategoryService interface {
	Create(claims ad.UserClaims, request dto.CreateCategoryRequest) error
	Update(claims ad.UserClaims, request dto.UpdateCategoryRequest, categoryId uint) error
	Delete(claims ad.UserClaims, categoryId uint) error
	GetAll(filter dto.FilterCategory) ([]dto.CategoryResponse, error)
	GetOne(categoryId uint) (dto.CategoryResponse, error)
}
