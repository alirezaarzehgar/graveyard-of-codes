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

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use: "consumer",
	Run: func(cmd *cobra.Command, args []string) {
		con, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Fatal(err)
		}
		defer con.Close()

		ch, err := con.Channel()
		if err != nil {
			log.Fatal(err)
		}
		defer ch.Close()

		err = ch.ExchangeDeclare("logs", amqp091.ExchangeFanout, true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		q, err := ch.QueueDeclare("", false, true, true, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		err = ch.QueueBind(q.Name, "", "logs", false, nil)
		if err != nil {
			log.Fatal(err)
		}

		msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		for m := range msgs {
			fmt.Println("Log:", string(m.Body))
			m.Ack(false)
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
