package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Println(string(msg.Subject), ":", string(msg.Data))
	})

	for {
		nc.Publish("foo", []byte("Hello, World"))
		time.Sleep(time.Second)
	}
}
