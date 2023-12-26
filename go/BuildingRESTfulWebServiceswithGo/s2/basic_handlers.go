package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
