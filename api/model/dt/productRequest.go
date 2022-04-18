package dt

import "go-restaurant-api/api/model/entity"

type ProductRequest interface {
	MapperToProduct() *entity.Product
}

type ProductCreateRequest struct {
	ProductName string  `json:"product_name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
	CategoryId  int64   `json:"category_id" validate:"required"`
}

type ProductUpdateRequest struct {
	Id int64 `json:"id" validate:"required"`
	ProductCreateRequest
}

func (p *ProductCreateRequest) MapperToProduct() *entity.Product {
	return &entity.Product{
		ProductName: p.ProductName,
		Description: p.Description,
		Price:       p.Price,
		Category: entity.Category{
			Id: p.CategoryId,
		},
	}
}

func (p *ProductUpdateRequest) MapperToProduct() *entity.Product {
	return &entity.Product{
		Id:          p.Id,
		ProductName: p.ProductName,
		Description: p.Description,
		Price:       p.Price,
		Category: entity.Category{
			Id: p.CategoryId,
		},
	}
}
