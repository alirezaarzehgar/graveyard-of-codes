package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Println("Subject:", msg.Subject)
		fmt.Println("Data:", string(msg.Data))
	})

	nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Println("Subject:", msg.Subject)
		fmt.Println("Data:", string(msg.Data))
	})

	nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Println("Subject:", msg.Subject)
		fmt.Println("Data:", string(msg.Data))
	})

	for i := 0; ; i++ {
		nc.Publish("foo", []byte(fmt.Sprint("Hello World ", i)))
		time.Sleep(time.Second * 2)
	}
}
