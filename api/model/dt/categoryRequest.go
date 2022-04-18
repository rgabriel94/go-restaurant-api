package dt

import "go-restaurant-api/api/model/entity"

type CategoryRequest interface {
	MapperToCategory() *entity.Category
}

type CategoryCreateRequest struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type CategoryUpdateRequest struct {
	Id int64 `json:"id" validate:"required"`
	CategoryCreateRequest
}

func (c *CategoryCreateRequest) MapperToCategory() *entity.Category {
	return &entity.Category{
		CategoryName: c.CategoryName,
	}
}

func (c *CategoryUpdateRequest) MapperToCategory() *entity.Category {
	return &entity.Category{
		Id:           c.Id,
		CategoryName: c.CategoryName,
	}
}
