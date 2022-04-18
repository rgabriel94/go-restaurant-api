package exception

import (
	"encoding/json"
	"go-restaurant-api/api/model/dt"
	"log"
	"net/http"
)

type ErrorException struct {
	StatusCode int
	Message    string `json:"message"`
}

func (exception *ErrorException) CreateHttpResponseException(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(exception.StatusCode)
	err := json.NewEncoder(w).Encode(dt.ErrorResponse{Message: exception.Message})
	if err != nil {
		log.Printf("Failed to write error message. Message - %s", exception.Message)
	}
}

func NewErrorResponse(statusCode int, message string) *ErrorException {
	return &ErrorException{
		StatusCode: statusCode,
		Message:    message,
	}
}

func PanicNotFound(message string) {
	panic(NewErrorResponse(http.StatusNotFound, message))
}

func PanicBadRequest(message string) {
	panic(NewErrorResponse(http.StatusBadRequest, message))
}
