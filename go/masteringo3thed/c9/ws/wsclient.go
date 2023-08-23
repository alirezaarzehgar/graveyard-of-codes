package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var in = bufio.NewReader(os.Stdin)

func getInput(input chan<- string) {
	fmt.Print("Enter message: ")
	text, err := in.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	input <- text
}

func main() {
	input := make(chan string, 1)
	go getInput(input)

	URL := url.URL{Scheme: "ws", Host: "localhost:8000", Path: "/ws"}
	ws, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, text, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("\rRecoeved message: %s\n", strings.TrimSpace(string(text)))
			fmt.Print("Enter message: ")
		}
	}()
	for {
		select {
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout!")
			return
		case <-done:
			return
		case t := <-input:
			err := ws.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println(err)
				return
			}
			go getInput(input)
		}
	}
}
