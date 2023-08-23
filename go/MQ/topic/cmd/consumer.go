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
	EXCHANGE_KEY = "logs_topic"
	AMQP_URL     = "amqp://guest:guest@localhost:5672/"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use: "consumer",
	Run: func(cmd *cobra.Command, args []string) {
		topics, _ := cmd.Flags().GetStringArray("topics")

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

		err = ch.ExchangeDeclare(EXCHANGE_KEY, amqp091.ExchangeTopic, true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		q, err := ch.QueueDeclare("", true, true, true, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		for _, topic := range topics {
			err = ch.QueueBind(q.Name, topic, EXCHANGE_KEY, false, nil)
			if err != nil {
				log.Fatal(err)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	consumerCmd.PersistentFlags().StringArray("topics", []string{"#"}, "Topics")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
