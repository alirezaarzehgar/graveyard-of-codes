/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Consume RabbitMQ messages",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatal(err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare("hiq", false, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		for msg := range msgs {
			fmt.Println(string(msg.Body))
		}
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
