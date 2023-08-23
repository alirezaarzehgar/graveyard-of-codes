package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		for {
			t, msg, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			fmt.Println("client:", strings.TrimSpace(string(msg)))

			err = ws.WriteMessage(t, msg)
			if err != nil {
				log.Println(err)
				break
			}
		}
	})

	server := http.Server{
		Addr:         ":8000",
		Handler:      mux,
		IdleTimeout:  time.Second * 3,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
	}
	server.ListenAndServe()
}
