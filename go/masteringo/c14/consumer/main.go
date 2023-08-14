package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Record struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     "justatopic",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	defer r.Close()
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("Message at offset %d: %v = %v\n", m.Offset, string(m.Key), string(m.Value))
		tmp := Record{}
		err = json.Unmarshal(m.Value, &tmp)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%T\n", tmp)
	}
}
