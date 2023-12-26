package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("OK"))
// 	})
// 	lg := handlers.LoggingHandler(os.Stdout, r)
// 	http.ListenAndServe(":8000", lg)
// }

func main() {
	http.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})))
	http.ListenAndServe(":8000", nil)
}
