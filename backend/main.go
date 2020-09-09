package main

import (
	"net/http"

	"etoneclab.com/basic-server/endpoints"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	muxRouter := mux.NewRouter()
	corsMw := mux.CORSMethodMiddleware(muxRouter)
	muxRouter.Use(corsMw)
	muxRouter.StrictSlash(true)

	endpoints.PrepareEndpoints(muxRouter)

	handl := handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "X-Token", "Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"}))(muxRouter)

	http.Handle("/", muxRouter)

	http.ListenAndServe(":8080", handl)

}
