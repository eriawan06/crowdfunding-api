package service

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/category/mapper"
	"crowdfunding-api/src/modules/category/model/dto"
	"crowdfunding-api/src/modules/category/repository"
	e "crowdfunding-api/src/utils/errors"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{Repository: repository}
}

func (service *CategoryServiceImpl) Create(claims ad.UserClaims, request dto.CreateCategoryRequest) error {
	if claims.Role != "admin" {
		return e.ErrForbidden
	}

	category := mapper.CreateCategoryRequestToCategory(claims, request)
	err := service.Repository.Create(category)
	return err
}

func (service *CategoryServiceImpl) Update(claims ad.UserClaims, request dto.UpdateCategoryRequest, categoryId uint) error {
	if claims.Role != "admin" {
		return e.ErrForbidden
	}

	// check
	_, err := service.Repository.FindOne(categoryId)
	if err != nil {
		return err
	}

	category := mapper.UpdateCategoryRequestToCategory(claims, request)
	err = service.Repository.Update(category, categoryId)
	return err
}

func (service *CategoryServiceImpl) Delete(claims ad.UserClaims, categoryId uint) error {
	if claims.Role != "admin" {
		return e.ErrForbidden
	}

	// check
	_, err := service.Repository.FindOne(categoryId)
	if err != nil {
		return err
	}

	err = service.Repository.Delete(categoryId)
	return err
}

func (service *CategoryServiceImpl) GetAll(filter dto.FilterCategory) ([]dto.CategoryResponse, error) {
	categories, err := service.Repository.FindAll(filter)
	if err != nil {
		return nil, err
	}

	response := mapper.ListCategoryToListCategoryResponse(categories)
	return response, nil
}

func (service *CategoryServiceImpl) GetOne(categoryId uint) (dto.CategoryResponse, error) {
	category, err := service.Repository.FindOne(categoryId)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	response := mapper.CategoryToCategoryResponse(category)
	return response, nil
}
