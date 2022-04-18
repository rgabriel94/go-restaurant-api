package middleware

import (
	"fmt"
	"go-restaurant-api/api/model/exception"
	"log"
	"net/http"
)

func CatchExceptionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer catchException(w)
		next.ServeHTTP(w, r)
	})
}

func catchException(w http.ResponseWriter) {
	err := recover()
	if err == nil {
		return
	}
	log.Println(err)
	switch err.(type) {
	case *exception.ErrorException:
		err.(*exception.ErrorException).CreateHttpResponseException(w)
	default:
		message := fmt.Sprintf("%s - We had a problem with our server", http.StatusText(http.StatusInternalServerError))
		exception.NewErrorResponse(http.StatusInternalServerError, message).CreateHttpResponseException(w)
	}
}
