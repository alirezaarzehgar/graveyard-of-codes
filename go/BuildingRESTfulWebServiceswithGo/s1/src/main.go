package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		if urlPathElements[1] == "romain_number" {
			number, _ := strconv.Atoi(urlPathElements[2])
			if number <= 0 || number > len(romainNumbers) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romainNumbers[number]))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})

	s := http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
