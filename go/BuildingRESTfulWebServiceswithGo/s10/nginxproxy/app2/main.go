package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "app2: %s\n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":5002", nil)
}
