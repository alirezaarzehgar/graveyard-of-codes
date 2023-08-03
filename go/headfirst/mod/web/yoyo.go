package web

import (
	"fmt"
	"net/http"
	"time"
)

func HelloWeb(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Wotrld WWW %s", time.Now())
	})
	http.ListenAndServe(addr, nil)
}
