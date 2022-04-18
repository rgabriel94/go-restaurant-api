package service

import (
	"go-restaurant-api/api/model/entity"
	"go-restaurant-api/api/repository"
)

var productRepository = repository.ProductRepo

func ListAllProducts() []entity.Product {
	return productRepository.ListAllProducts()
}

func GetProduct(productId int64) *entity.Product {
	return productRepository.GetProduct(productId)
}

func CreateProduct(product *entity.Product) {
	productRepository.CreateProduct(product)
	category := GetCategory(product.Category.Id)
	product.Category.CategoryName = category.CategoryName
}

func UpdateProduct(product *entity.Product) *entity.Product {
	productRepository.UpdateProduct(product)
	return GetProduct(product.Id)
}

func DeleteProduct(productId int64) {
	productRepository.DeleteProduct(productId)
}
