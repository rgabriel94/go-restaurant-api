package service

import (
	"go-restaurant-api/api/model/entity"
	"go-restaurant-api/api/repository"
)

var categoryRepository = repository.CategoryRepo

func ListAllCategories() []entity.Category {
	return categoryRepository.ListAllCategories()
}

func GetCategory(categoryId int64) *entity.Category {
	return categoryRepository.GetCategory(categoryId)
}

func CreateCategory(category *entity.Category) {
	categoryRepository.CreateCategory(category)
}

func UpdateCategory(category *entity.Category) {
	categoryRepository.UpdateCategory(category)
}

func DeleteCategory(categoryId int64) {
	categoryRepository.DeleteCategory(categoryId)
}
