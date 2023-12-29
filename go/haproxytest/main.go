package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync/atomic"
)

func main() {
	port := flag.Uint("port", 5001, "port number")
	flag.Parse()

	var counter int64
	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&counter, 1)
	})
	http.HandleFunc("/i", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "count: %d\n", counter)
	})

	http.ListenAndServe(fmt.Sprint(":", *port), nil)
}
