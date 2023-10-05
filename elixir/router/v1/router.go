package router

import (
	"net/http"

	"github.com/amagimedia/elixir/elixir/handler/v1"

	"github.com/gorilla/mux"
)

func AttachRoutes(router *mux.Router) {
	// attach the routes here
	router.HandleFunc("/api/v1/hello", handler.Bar).Methods(http.MethodGet)
}
