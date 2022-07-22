package dto

type CreateCategoryRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type UpdateCategoryRequest struct {
	CreateCategoryRequest
}

type CategoryResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type FilterCategory struct {
	Name     string
	IsActive string
}
