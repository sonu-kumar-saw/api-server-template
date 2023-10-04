package router

import (
	"api-server-template/cmd/handler/v1"
	"net/http"

	"github.com/gorilla/mux"
)

func AttachRoutes(router *mux.Router) {
	// attach the routes here
	router.HandleFunc("/api/v1/hello", handler.Bar).Methods(http.MethodGet)
}
