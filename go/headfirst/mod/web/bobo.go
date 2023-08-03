package web

import (
	"fmt"
	"net/http"
	"time"
)

func Bobo(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bo bo !!! %s", time.Now())
	})
	http.ListenAndServe(addr, nil)
}
