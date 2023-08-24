/*
Copyright © 2023 Alireza Arzehgar
*/
package cmd

import (
	"context"
	"log"
	"orders/conf"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

// reserveCmd represents the reserve command
var reserveCmd = &cobra.Command{
	Use: "reserve",
	Run: func(cmd *cobra.Command, args []string) {
		topic, _ := cmd.Flags().GetString("topic")

		ch, err := conf.DialAmqp()
		if err != nil {
			log.Fatal(err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(topic, true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		err = ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Some product"),
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(reserveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	reserveCmd.PersistentFlags().String("topic", "", "get topic")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reserveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
