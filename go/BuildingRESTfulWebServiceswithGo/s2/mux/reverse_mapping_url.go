package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hi/{content}", func(w http.ResponseWriter, r *http.Request) {}).Name("Hi")

	u, err := r.Get("Hi").URL("content", "hoyya")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
}
