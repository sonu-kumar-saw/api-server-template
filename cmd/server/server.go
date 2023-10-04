package server

import (
	v1router "api-server-template/cmd/router/v1"
	v2router "api-server-template/cmd/router/v2"
	"net/http"

	"github.com/gorilla/mux"
)

// implement the server here

func Run() {

	// create new router
	router := mux.NewRouter()
	// attach v1 router
	v1router.AttachRoutes(router)
	// attach v2 router
	v2router.AttachRoutes(router)
	// start listening
	http.ListenAndServe(":8080", router)

}
