/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

const (
	EXCHANGE_NAME = "logs_routing"
	AMQP_URL      = "amqp://guest:guest@localhost:5672/"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use: "consumer",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := amqp091.Dial(AMQP_URL)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatal(err)
		}
		defer ch.Close()

		err = ch.ExchangeDeclare(EXCHANGE_NAME, amqp091.ExchangeDirect, true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		q, err := ch.QueueDeclare("", true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		for _, ll := range []string{"debug", "info", "critical", "emergency"} {
			fmt.Println("Binding queue", q.Name, "to exchange logs with routing key", ll)
			err = ch.QueueBind(q.Name, ll, EXCHANGE_NAME, false, nil)
			if err != nil {
				log.Println(err)
			}
		}

		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		for m := range msgs {
			fmt.Println(m.RoutingKey, "=>", string(m.Body))
		}
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
