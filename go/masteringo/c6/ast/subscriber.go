package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	_, _ = nc.Subscribe("foo", func(m *nats.Msg) {
		log.Printf("Received a message on %s: %s\n", m.Subject, string(m.Data))
	})

	select {}
}
