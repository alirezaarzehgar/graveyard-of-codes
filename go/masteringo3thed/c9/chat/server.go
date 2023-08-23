// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{}

// type SocketReader struct {
// 	c    *websocket.Conn
// 	mode int
// 	name string
// }

// func (i *SocketReader) broadcast(str string) {
// 	for _, reader := range socketReaders {
// 		if reader == i || reader.mode == 1 {
// 			continue
// 		}
// 		reader.writeMsg(i.name, str)
// 	}
// }

// func (i *SocketReader) read() {
// 	_, msg, err := i.c.ReadMessage()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(i.name, ":", string(msg))
// 	fmt.Println(i.mode)

// 	if i.mode == 1 {
// 		i.name = strings.TrimSpace(string(msg))
// 		i.writeMsg("System", "Welcome "+i.name+", please write a message and we will broadcast it to other users.")
// 		i.mode = 2

// 		return
// 	}

// 	i.broadcast(string(msg))
// 	fmt.Println(i.name, ":", string(msg))
// }

// func (i *SocketReader) startThread() {
// 	i.writeMsg("System", "Please write your name")
// 	i.mode = 1

// 	go func() {
// 		for {
// 			i.read()
// 		}
// 	}()
// }

// func (i *SocketReader) writeMsg(name, msg string) {
// 	i.c.WriteMessage(websocket.TextMessage, []byte(name+": "+msg))
// }

// var socketReaders []*SocketReader

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	if socketReaders == nil {
// 		socketReaders = make([]*SocketReader, 0)
// 	}

// 	ws, _ := upgrader.Upgrade(w, r, nil)
// 	ptrSock := &SocketReader{c: ws}
// 	socketReaders = append(socketReaders, ptrSock)
// 	ptrSock.startThread()
// }

// func main() {
// 	http.HandleFunc("/ws", wsHandler)
// 	http.ListenAndServe(":8000", nil)
// }
