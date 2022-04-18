package testutil

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go-restaurant-api/api/model/dt"
	"testing"
)

func NewCategoryCreateRequest() dt.CategoryCreateRequest {
	return dt.CategoryCreateRequest{
		CategoryName: RandomName(),
	}
}

func NewCategoryUpdateRequest(id int64) dt.CategoryUpdateRequest {
	return dt.CategoryUpdateRequest{
		Id:                    id,
		CategoryCreateRequest: NewCategoryCreateRequest(),
	}
}

func DecodeCategoryResponse(t *testing.T, body *bytes.Buffer) *dt.CategoryResponse {
	var categoryResponse dt.CategoryResponse
	err := json.NewDecoder(body).Decode(&categoryResponse)
	require.NoError(t, err)
	return &categoryResponse
}

func DecodeCategoriesResponse(t *testing.T, body *bytes.Buffer) []dt.CategoryResponse {
	var categoryResponse []dt.CategoryResponse
	err := json.NewDecoder(body).Decode(&categoryResponse)
	require.NoError(t, err)
	return categoryResponse
}
