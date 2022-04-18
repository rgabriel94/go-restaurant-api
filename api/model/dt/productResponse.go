package dt

import (
	"go-restaurant-api/api/model/entity"
	"time"
)

type ProductResponse struct {
	Id          int64            `json:"id"`
	ProductName string           `json:"product_name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	CreatedAt   time.Time        `json:"created_at"`
	Category    CategoryResponse `json:"category_id"`
}

func MapperToProductResponse(product *entity.Product) *ProductResponse {
	return &ProductResponse{
		Id:          product.Id,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		Category:    *MapperToCategoryResponse(&product.Category),
	}
}

func MapperToProductsResponse(products []entity.Product) []ProductResponse {
	productsResponse := make([]ProductResponse, len(products))
	for i := 0; i < len(products); i++ {
		productsResponse[i] = *MapperToProductResponse(&products[i])
	}
	return productsResponse
}
