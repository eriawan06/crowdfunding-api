package repository

import (
	"crowdfunding-api/src/modules/category/model/dto"
	"crowdfunding-api/src/modules/category/model/entity"
)

type CategoryRepository interface {
	Create(category entity.Category) error
	Update(category entity.Category, categoryId uint) error
	Delete(categoryId uint) error
	FindAll(filter dto.FilterCategory) ([]entity.Category, error)
	FindOne(categoryId uint) (entity.Category, error)
}
