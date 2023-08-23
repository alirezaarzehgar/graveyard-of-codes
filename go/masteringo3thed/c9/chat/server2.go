package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	var upgrader = websocket.Upgrader{}
	var clients = make([]*websocket.Conn, 0)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, _ := upgrader.Upgrade(w, r, nil)
		clients = append(clients, ws)

		go func() {
			for {
				_, msg, err := ws.ReadMessage()
				if err != nil {
					log.Println(err)
					ws.Close()
					return
				}

				for _, c := range clients {
					if c == ws {
						continue
					}
					c.WriteMessage(websocket.TextMessage, msg)
				}
			}
		}()
	})
	http.ListenAndServe(":8000", nil)
}
