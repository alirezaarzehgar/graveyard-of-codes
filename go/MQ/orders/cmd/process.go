/*
Copyright © 2023 Alireza Arzehgar
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"orders/conf"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use: "process",
	Run: func(cmd *cobra.Command, args []string) {
		topic, _ := cmd.Flags().GetString("topic")

		ch, err := conf.DialAmqp()
		if err != nil {
			log.Fatal(err)
		}
		defer ch.Close()

		msgs, err := ch.Consume(topic, "", false, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		C := time.After(time.Second / 2)
		for {
			select {
			case <-C:
				fmt.Print("Do you want to wait ? [Y/n]")
				out, err := bufio.NewReader(os.Stdin).ReadString('\n')
				if err == nil && strings.ToUpper(strings.TrimSpace(out)) == "Y" {
					fmt.Println("Wait for queue")
				} else {
					return
				}
			case m := <-msgs:
				fmt.Println(string(m.Body))
				fmt.Println(m.RoutingKey)

				fmt.Print("Done [Y/n]")
				out, err := bufio.NewReader(os.Stdin).ReadString('\n')
				if err == nil && strings.ToUpper(strings.TrimSpace(out)) == "Y" {
					m.Ack(false)
					fmt.Println("Done")
				} else {
					m.Nack(false, true)
					fmt.Println("Cancel")
					return
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(processCmd)
	processCmd.PersistentFlags().String("topic", "", "get topic")
}
