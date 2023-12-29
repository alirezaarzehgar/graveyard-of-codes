package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "app3: %d\n", rand.Intn(10))
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":5003", nil)
}
