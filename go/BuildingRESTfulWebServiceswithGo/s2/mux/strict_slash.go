package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/hi/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello, World!") })
	http.ListenAndServe(":8000", r)
}
