package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-restaurant-api/api/model/exception"
	"log"
	"net/http"
	"strconv"
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"

	base    = 10
	bitSize = 64
)

func ResponseWriter(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(statusCode)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}

func ExtractIntVars(r *http.Request, paramKey string) int64 {
	param := mux.Vars(r)[paramKey]
	value, err := strconv.ParseInt(param, base, bitSize)
	if err != nil {
		message := fmt.Sprintf("Method %s for path %s. Param %s invalid.",
			r.Method,
			r.RequestURI,
			param)
		log.Println(message)
		exception.PanicBadRequest(message)
	}
	return value
}
