package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	nc.Publish("foo", []byte("Hello World"))
	log.Println("Published message on subject 'foo'")
	nc.Close()
}
