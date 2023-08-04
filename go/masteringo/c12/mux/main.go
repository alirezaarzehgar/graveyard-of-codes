package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Root!")
	})
	http.ListenAndServe(":8000", r)
}
