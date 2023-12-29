package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		fmt.Fprintf(w, "app1: %s\n", string(body))
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":5001", nil)
}
