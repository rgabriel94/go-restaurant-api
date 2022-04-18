package controller

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go-restaurant-api/api/config/server"
	"go-restaurant-api/api/model/dt"
	"go-restaurant-api/test/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	productsUrl  = "/products"
	productIdUrl = "/products/%d"
)

func getProduct(t *testing.T, id int64) *dt.ProductResponse {
	url := fmt.Sprintf(productIdUrl, id)
	response := testutil.HttpServer(t, http.MethodGet, url, nil, http.StatusOK)
	return testutil.DecodeProductResponse(t, response.Body)
}

func createRandomProduct(t *testing.T, categoryId int64) *dt.ProductResponse {
	productRequest := testutil.NewProductCreateRequest(categoryId)
	body := testutil.PrepareBody(t, productRequest)
	response := testutil.HttpServer(t, http.MethodPost, productsUrl, body, http.StatusCreated)
	productResponse := testutil.DecodeProductResponse(t, response.Body)

	require.NotEqual(t, 0, productResponse.Id)
	require.Equal(t, productRequest.ProductName, productResponse.ProductName)
	require.Equal(t, productRequest.Description, productResponse.Description)
	require.Equal(t, productRequest.Price, productResponse.Price)
	require.NotEmpty(t, productResponse.CreatedAt)
	require.Equal(t, productRequest.CategoryId, productResponse.Category.Id)
	require.NotEmpty(t, productResponse.Category.CategoryName)

	return productResponse
}

func deleteProduct(t *testing.T, id int64) {
	url := fmt.Sprintf(productIdUrl, id)
	testutil.HttpServer(t, http.MethodDelete, url, nil, http.StatusNoContent)
}

func TestGetProduct(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	productResponse := createRandomProduct(t, categoryResponse.Id)
	productGetResponse := getProduct(t, productResponse.Id)
	require.Equal(t, productResponse.Id, productGetResponse.Id)
	require.Equal(t, productResponse.ProductName, productGetResponse.ProductName)
	require.Equal(t, productResponse.Description, productGetResponse.Description)
	require.Equal(t, productResponse.Price, productGetResponse.Price)
	require.NotEmpty(t, productResponse.CreatedAt)
	require.Equal(t, productResponse.Category.Id, productGetResponse.Category.Id)
	require.Equal(t, productResponse.Category.CategoryName, productGetResponse.Category.CategoryName)
	deleteProduct(t, productResponse.Id)
	deleteCategory(t, categoryResponse.Id)
}

func TestCreateProductHandler(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	productResponse := createRandomProduct(t, categoryResponse.Id)
	deleteProduct(t, productResponse.Id)
	deleteCategory(t, categoryResponse.Id)
}

func TestUpdateProductHandler(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	categoryForUpdateResponse := createRandomCategory(t)
	productCreateResponse := createRandomProduct(t, categoryResponse.Id)

	productUpdateRequest := testutil.NewProductUpdateRequest(productCreateResponse.Id, categoryForUpdateResponse.Id)
	body := testutil.PrepareBody(t, productUpdateRequest)
	response := testutil.HttpServer(t, http.MethodPut, productsUrl, body, http.StatusOK)
	productUpdateResponse := testutil.DecodeProductResponse(t, response.Body)

	require.NotEqual(t, productUpdateRequest, productUpdateResponse.Id)
	require.Equal(t, productUpdateRequest.ProductName, productUpdateResponse.ProductName)
	require.Equal(t, productUpdateRequest.Description, productUpdateResponse.Description)
	require.Equal(t, productUpdateRequest.Price, productUpdateResponse.Price)
	require.NotEmpty(t, productUpdateResponse.CreatedAt)
	require.Equal(t, categoryForUpdateResponse.Id, productUpdateResponse.Category.Id)
	require.Equal(t, categoryForUpdateResponse.CategoryName, productUpdateResponse.Category.CategoryName)

	deleteProduct(t, productUpdateResponse.Id)
	deleteCategory(t, categoryForUpdateResponse.Id)
	deleteCategory(t, categoryResponse.Id)
}

func TestDeleteProductHandler(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	productResponse := createRandomProduct(t, categoryResponse.Id)

	deleteProduct(t, productResponse.Id)
	url := fmt.Sprintf(productIdUrl, productResponse.Id)
	request := httptest.NewRequest(http.MethodGet, url, nil)
	response := httptest.NewRecorder()
	server.Router.ServeHTTP(response, request)
	require.Equal(t, http.StatusNotFound, response.Code)

	deleteCategory(t, categoryResponse.Id)
}
