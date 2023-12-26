package main

import (
	"fmt"
	"net/http"
)

func mid1(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before mid1")
		handler.ServeHTTP(w, r)
		fmt.Println("After mid1")
	})
}

func mid2(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before mid1")
		handler.ServeHTTP(w, r)
		fmt.Println("After mid1")
	})
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("On handler")
		fmt.Fprintln(w, "OK")
	})
	http.Handle("/r1", mid1(handler))
	http.HandleFunc("/r2", mid2(handler))
	http.ListenAndServe(":8000", nil)
}
