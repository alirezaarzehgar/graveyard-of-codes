/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use: "producer",
	Run: func(cmd *cobra.Command, args []string) {
		topic, _ := cmd.Flags().GetString("topic")
		logMsg, _ := cmd.Flags().GetString("log")

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

		ch.PublishWithContext(context.Background(), EXCHANGE_KEY, topic, false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(logMsg),
		})
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	producerCmd.PersistentFlags().String("topic", "#", "topic")
	producerCmd.PersistentFlags().String("log", "Hi!", "log message")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// producerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
