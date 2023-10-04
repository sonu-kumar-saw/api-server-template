package router

import (
	"api-server-template/cmd/handler/v2"
	"net/http"

	"github.com/gorilla/mux"
)

func AttachRoutes(router *mux.Router) {
	// attach the routes here
	router.HandleFunc("/api/v2/hello", handler.Foo).Methods(http.MethodGet)
}
