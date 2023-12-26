package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		fmt.Fprintln(w, q.Get("a"), q.Get("b"), q.Get("c"))
	})
	http.ListenAndServe(":8000", r)
}
