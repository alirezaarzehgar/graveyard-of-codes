package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func articleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "category: %s\n", vars["category"])
	fmt.Fprintf(w, "id: %s\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", articleHandler)
	srv := http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
