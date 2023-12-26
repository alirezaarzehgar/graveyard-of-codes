package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	lip := r.Host("127.0.0.1").Subrouter()
	loc := r.Host("localhost").Subrouter()

	lip.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello, 127.0.0.1!") })
	loc.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello, localhost!") })
	http.ListenAndServe(":8000", r)
}
