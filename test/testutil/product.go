package testutil

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go-restaurant-api/api/model/dt"
	"testing"
)

func NewProductCreateRequest(categoryId int64) dt.ProductCreateRequest {
	return dt.ProductCreateRequest{
		ProductName: RandomName(),
		Description: RandomString(150),
		Price:       RandomFloat(),
		CategoryId:  categoryId,
	}
}

func NewProductUpdateRequest(productId int64, categoryId int64) dt.ProductUpdateRequest {
	return dt.ProductUpdateRequest{
		Id:                   productId,
		ProductCreateRequest: NewProductCreateRequest(categoryId),
	}
}

func DecodeProductResponse(t *testing.T, body *bytes.Buffer) *dt.ProductResponse {
	var productResponse dt.ProductResponse
	err := json.NewDecoder(body).Decode(&productResponse)
	require.NoError(t, err)
	return &productResponse
}
