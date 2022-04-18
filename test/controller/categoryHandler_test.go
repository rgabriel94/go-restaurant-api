package controller

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go-restaurant-api/api/config/server"
	"go-restaurant-api/api/model/dt"
	"go-restaurant-api/test/testutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	categoriesUrl = "/categories"
	categoryIdUrl = "/categories/%d"
)

func getCategory(t *testing.T, id int64) *dt.CategoryResponse {
	url := fmt.Sprintf(categoryIdUrl, id)
	response := testutil.HttpServer(t, http.MethodGet, url, nil, http.StatusOK)
	return testutil.DecodeCategoryResponse(t, response.Body)
}

func createRandomCategory(t *testing.T) *dt.CategoryResponse {
	categoryRequest := testutil.NewCategoryCreateRequest()
	body := testutil.PrepareBody(t, categoryRequest)
	response := testutil.HttpServer(t, http.MethodPost, categoriesUrl, body, http.StatusCreated)
	categoryResponse := testutil.DecodeCategoryResponse(t, response.Body)
	require.NotEqual(t, 0, categoryResponse.Id)
	require.Equal(t, categoryRequest.CategoryName, categoryResponse.CategoryName)
	return categoryResponse
}

func deleteCategory(t *testing.T, id int64) {
	url := fmt.Sprintf(categoryIdUrl, id)
	testutil.HttpServer(t, http.MethodDelete, url, nil, http.StatusNoContent)
}

func TestListAllCategories(t *testing.T) {
	categoriesCreateResponse := make([]dt.CategoryResponse, 2)
	categoriesCreateResponse[0] = *createRandomCategory(t)
	categoriesCreateResponse[1] = *createRandomCategory(t)
	response := testutil.HttpServer(t, http.MethodGet, categoriesUrl, nil, http.StatusOK)
	categoriesGetResponse := testutil.DecodeCategoriesResponse(t, response.Body)

	sizeCreate := len(categoriesCreateResponse)
	sizeGet := len(categoriesGetResponse)

	log.Println(sizeCreate)
	log.Println(sizeGet)
	require.GreaterOrEqual(t, sizeGet, sizeCreate)
	i, j := sizeGet-1, sizeCreate-1
	for ; i >= 0; i-- {
		if j < 0 {
			break
		}
		if categoriesCreateResponse[j].Id == categoriesGetResponse[i].Id {
			require.Equal(t, categoriesCreateResponse[j].CategoryName, categoriesGetResponse[i].CategoryName)
			j--
		}
	}
	if j >= 0 {
		t.Errorf("The created categories were not found.")
	}
	deleteCategory(t, categoriesCreateResponse[0].Id)
	deleteCategory(t, categoriesCreateResponse[1].Id)
}

func TestGetCategory(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	categoryGetResponse := getCategory(t, categoryResponse.Id)
	require.Equal(t, categoryResponse.Id, categoryGetResponse.Id)
	require.Equal(t, categoryResponse.CategoryName, categoryGetResponse.CategoryName)
	deleteCategory(t, categoryResponse.Id)
}

func TestCreateCategoryHandler(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	deleteCategory(t, categoryResponse.Id)
}

func TestUpdateCategoryHandler(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	categoryUpdateRequest := testutil.NewCategoryUpdateRequest(categoryResponse.Id)

	body := testutil.PrepareBody(t, categoryUpdateRequest)
	response := testutil.HttpServer(t, http.MethodPut, categoriesUrl, body, http.StatusOK)
	categoryUpdateResponse := testutil.DecodeCategoryResponse(t, response.Body)
	require.Equal(t, categoryResponse.Id, categoryUpdateResponse.Id)
	require.Equal(t, categoryUpdateRequest.CategoryName, categoryUpdateResponse.CategoryName)
	deleteCategory(t, categoryResponse.Id)
}

func TestDeleteCategoryHandler(t *testing.T) {
	categoryResponse := createRandomCategory(t)
	deleteCategory(t, categoryResponse.Id)
	url := fmt.Sprintf(categoryIdUrl, categoryResponse.Id)
	request := httptest.NewRequest(http.MethodGet, url, nil)
	response := httptest.NewRecorder()
	server.Router.ServeHTTP(response, request)
	require.Equal(t, http.StatusNotFound, response.Code)
}
