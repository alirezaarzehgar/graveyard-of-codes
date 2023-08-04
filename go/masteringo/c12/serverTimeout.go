package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, time.Now().Format(time.RFC1123))
	})
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "path:", r.URL.Path)
		fmt.Fprintln(w, "host:", r.URL.Host)
	})

	server := http.Server{
		Addr:         ":8000",
		Handler:      m,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
	}
	server.ListenAndServe()
}
