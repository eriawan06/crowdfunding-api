package mapper

import (
	ad "crowdfunding-api/src/modules/auth/model/dto"
	"crowdfunding-api/src/modules/category/model/dto"
	"crowdfunding-api/src/modules/category/model/entity"
	"crowdfunding-api/src/utils/common"
)

func CreateCategoryRequestToCategory(claims ad.UserClaims, request dto.CreateCategoryRequest) entity.Category {
	return entity.Category{
		Name:     request.Name,
		IsActive: &request.IsActive,
		BaseEntity: common.BaseEntity{
			CreatedBy: claims.Email,
			UpdatedBy: claims.Email,
		},
	}
}

func UpdateCategoryRequestToCategory(claims ad.UserClaims, request dto.UpdateCategoryRequest) entity.Category {
	return entity.Category{
		Name:     request.Name,
		IsActive: &request.IsActive,
		BaseEntity: common.BaseEntity{
			UpdatedBy: claims.Email,
		},
	}
}

func CategoryToCategoryResponse(category entity.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id:       category.Id,
		Name:     category.Name,
		IsActive: *category.IsActive,
	}
}

func ListCategoryToListCategoryResponse(categories []entity.Category) []dto.CategoryResponse {
	var categoriesResp []dto.CategoryResponse
	for _, v := range categories {
		categoriesResp = append(categoriesResp, CategoryToCategoryResponse(v))
	}
	return categoriesResp
}
