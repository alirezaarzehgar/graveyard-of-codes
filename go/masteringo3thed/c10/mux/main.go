package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/regex/{value:be+p}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("path:", r.URL.Path)
	})
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed\n")
	})
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Header.Get("user")
			if user == "admin" {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	})
	admin.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("user"))
		fmt.Fprintf(w, "Hi admin!\n")
	}).Methods(http.MethodGet)
	admin.Handle("/oh", http.StripPrefix("/admin", http.FileServer(http.Dir("."))))

	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	go func() {
		s.ListenAndServe()
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	fmt.Println("Quite after signal:", <-sigs)
	s.Shutdown(context.TODO())
}
