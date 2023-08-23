/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Produce RabbitMQ messages",
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

		err = ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello fucking world"),
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// producerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// producerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
