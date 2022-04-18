package testutil

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go-restaurant-api/api/config/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HttpServer(t *testing.T, method string, url string, body *bytes.Reader, statusCode int) *httptest.ResponseRecorder {
	var request *http.Request
	if body != nil {
		request = httptest.NewRequest(method, url, body)
	} else {
		request = httptest.NewRequest(method, url, nil)
	}
	response := httptest.NewRecorder()
	server.Router.ServeHTTP(response, request)
	require.Equal(t, statusCode, response.Code)
	return response
}

func PrepareBody(t *testing.T, preBody any) *bytes.Reader {
	body, err := json.Marshal(preBody)
	require.NoError(t, err)
	return bytes.NewReader(body)
}
