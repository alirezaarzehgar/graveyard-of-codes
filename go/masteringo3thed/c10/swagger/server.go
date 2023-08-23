package main

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/docs", middleware.Redoc(middleware.RedocOpts{SpecURL: "./swagger.yaml"}, nil)).Methods(http.MethodGet)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", r)

}
