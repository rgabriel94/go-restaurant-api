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
	productsPath   = "/products"
	paramProductId = "product_id"
)

func ProductHandlerConfiguration(router *mux.Router) {
	log.Println("Adding product handlers.")
	router.HandleFunc(productsPath, listAllProducts).Methods(http.MethodGet)
	router.HandleFunc(productsPath+"/{"+paramProductId+"}", getProduct).Methods(http.MethodGet)
	router.HandleFunc(productsPath, createProduct).Methods(http.MethodPost)
	router.HandleFunc(productsPath, updateProduct).Methods(http.MethodPut)
	router.HandleFunc(productsPath+"/{"+paramProductId+"}", deleteProduct).Methods(http.MethodDelete)
	log.Println("Added product handlers.")
}

func listAllProducts(w http.ResponseWriter, r *http.Request) {
	products := service.ListAllProducts()
	productsResponse := dt.MapperToProductsResponse(products)
	service.ResponseWriter(w, http.StatusOK, productsResponse)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	productId := service.ExtractIntVars(r, paramProductId)
	product := service.GetProduct(productId)
	productResponse := dt.MapperToProductResponse(product)
	service.ResponseWriter(w, http.StatusOK, productResponse)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	product := extractProductFromTheBody(r, &dt.ProductCreateRequest{})
	service.CreateProduct(product)
	productResponse := dt.MapperToProductResponse(product)
	service.ResponseWriter(w, http.StatusCreated, productResponse)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	product := extractProductFromTheBody(r, &dt.ProductUpdateRequest{})
	product = service.UpdateProduct(product)
	productResponse := dt.MapperToProductResponse(product)
	service.ResponseWriter(w, http.StatusOK, productResponse)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := service.ExtractIntVars(r, paramProductId)
	service.DeleteProduct(productId)
	service.ResponseWriter(w, http.StatusNoContent, nil)
}

func extractProductFromTheBody(r *http.Request, productRequest dt.ProductRequest) *entity.Product {
	extractProductRequestFromTheBody(r, productRequest)
	service.Validate(productRequest)
	return productRequest.MapperToProduct()
}

func extractProductRequestFromTheBody(r *http.Request, productRequest dt.ProductRequest) {
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		log.Println(err)
		exception.PanicBadRequest(enum.ProductBodyExpected)
	}
}
