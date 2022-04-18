package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-restaurant-api/api/enum"
	"go-restaurant-api/api/model/dt"
	"go-restaurant-api/api/model/entity"
	"go-restaurant-api/api/model/exception"
	"go-restaurant-api/api/service"
	"log"
	"net/http"
)

const (
	paramCategoryId = "category_id"
	categoriesPath  = "/categories"
)

func CategoryHandlerConfiguration(router *mux.Router) {
	log.Println("Adding category handlers.")
	router.HandleFunc(categoriesPath, listAllCategories).Methods(http.MethodGet)
	router.HandleFunc(categoriesPath+"/{"+paramCategoryId+"}", getCategory).Methods(http.MethodGet)
	router.HandleFunc(categoriesPath, createCategory).Methods(http.MethodPost)
	router.HandleFunc(categoriesPath, updateCategory).Methods(http.MethodPut)
	router.HandleFunc(categoriesPath+"/{"+paramCategoryId+"}", deleteCategory).Methods(http.MethodDelete)
	log.Println("Added category handlers.")
}

func listAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := service.ListAllCategories()
	categoryResponse := dt.MapperToCategoriesResponse(categories)
	service.ResponseWriter(w, http.StatusOK, categoryResponse)
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := service.ExtractIntVars(r, paramCategoryId)
	category := service.GetCategory(categoryId)
	categoryResponse := dt.MapperToCategoryResponse(category)
	service.ResponseWriter(w, http.StatusOK, categoryResponse)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	category := extractCategoryFromTheBody(r, &dt.CategoryCreateRequest{})
	service.CreateCategory(category)
	categoryResponse := dt.MapperToCategoryResponse(category)
	service.ResponseWriter(w, http.StatusCreated, categoryResponse)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	category := extractCategoryFromTheBody(r, &dt.CategoryUpdateRequest{})
	service.UpdateCategory(category)
	categoryResponse := dt.MapperToCategoryResponse(category)
	service.ResponseWriter(w, http.StatusOK, categoryResponse)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := service.ExtractIntVars(r, paramCategoryId)
	service.DeleteCategory(categoryId)
	service.ResponseWriter(w, http.StatusNoContent, nil)
}

func extractCategoryFromTheBody(r *http.Request, categoryRequest dt.CategoryRequest) *entity.Category {
	extractCategoryRequestFromTheBody(r, categoryRequest)
	service.Validate(categoryRequest)
	return categoryRequest.MapperToCategory()
}

func extractCategoryRequestFromTheBody(r *http.Request, categoryRequest dt.CategoryRequest) {
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		log.Println(err)
		exception.PanicBadRequest(enum.CategoryBodyExpected)
	}
}
