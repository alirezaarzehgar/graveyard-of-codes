/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use: "producer",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := cmd.Flags().GetStringArray("tasks")
		fmt.Println(tasks)

		conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatal(err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare("hiq", true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		for _, task := range tasks {
			ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp091.Publishing{
				DeliveryMode: amqp091.Persistent,
				ContentType:  "plain/text",
				Body:         []byte(task),
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	producerCmd.PersistentFlags().StringArray("tasks", []string{"Task."}, "Get tasks")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// producerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
