package server

import (
	"github.com/gorilla/mux"
	"go-restaurant-api/api/controller"
	"go-restaurant-api/api/middleware"
	"log"
	"net/http"
	"strconv"
)

var Router *mux.Router

func init() {
	log.Println("Creating router...")
	Router = mux.NewRouter().StrictSlash(true)
	addMiddlewares()
	addHandlers()
	log.Println("Router created.")
}

func addMiddlewares() {
	Router.Use(middleware.CatchExceptionMiddleware)
}

func addHandlers() {
	controller.CategoryHandlerConfiguration(Router)
	controller.ProductHandlerConfiguration(Router)
}

func StartServer(port int) {
	log.Println("Starting server...")
	if port == 0 {
		port = 8080
	}
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(port), Router))
}
