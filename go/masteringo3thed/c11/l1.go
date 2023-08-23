package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

//go:generate echo hi generation!
//go:generate lsmem
func main() {
	r := mux.NewRouter()
	getMux := r.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/hi", handler)
	getMux.HandleFunc("/hyo", handler)
	getMux.HandleFunc("/user/{id:[0-9]+}", handler)

	postMux := r.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/hi", handler)
	postMux.HandleFunc("/hyo", handler)
	postMux.HandleFunc("/user/{id:[0-9]+}", handler)

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		method, err := route.GetMethods()
		if err == nil {
			fmt.Print("|METHOD:", method)
		}

		pt, err := route.GetPathTemplate()
		if err == nil {
			fmt.Print("|PATHTEMPLATE", pt)
		}

		pr, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("|PATHREGEXP:", pr)
		}

		return nil
	})

	http.ListenAndServe(":8000", r)
}
