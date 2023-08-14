package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "justatopic", 0)
	if err != nil {
		log.Fatal("DialLeader:", err)
	}
	defer conn.Close()

	for i := 0; i < 15; i++ {
		data := struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{Name: fmt.Sprint("Name ", i), Age: i * 5}

		raw, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			continue
		}
		conn.WriteMessages(kafka.Message{Value: raw})
	}
}
