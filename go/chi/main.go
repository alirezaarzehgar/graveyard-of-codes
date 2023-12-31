package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!" + chi.URLParam(r, "id")))
	})
	http.ListenAndServe(":3000", r)
}
