package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alirezaarzehgar/romainserver/mymux"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		for s := range sig {
			log.Println("SIGINT captured: ", s)
		}
	}()

	mux := mymux.New()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World")
	})
	log.Fatal(http.ListenAndServe(":8000", mux))
}
