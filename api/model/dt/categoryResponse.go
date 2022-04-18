package dt

import (
	"go-restaurant-api/api/model/entity"
)

type CategoryResponse struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"`
}

func MapperToCategoryResponse(category *entity.Category) *CategoryResponse {
	return &CategoryResponse{
		Id:           category.Id,
		CategoryName: category.CategoryName,
	}
}

func MapperToCategoriesResponse(categories []entity.Category) []CategoryResponse {
	categoryResponse := make([]CategoryResponse, len(categories))
	for i := 0; i < len(categories); i++ {
		categoryResponse[i] = *MapperToCategoryResponse(&categories[i])
	}
	return categoryResponse
}
