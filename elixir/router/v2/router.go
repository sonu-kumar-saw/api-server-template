package router

import (
	"net/http"

	"github.com/amagimedia/elixir/elixir/handler/v2"

	"github.com/gorilla/mux"
)

func AttachRoutes(router *mux.Router) {
	// attach the routes here
	router.HandleFunc("/api/v2/hello", handler.Foo).Methods(http.MethodGet)
}
