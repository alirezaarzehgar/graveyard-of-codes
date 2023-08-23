/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Connect to a unix socket server",
	Run: func(cmd *cobra.Command, args []string) {
		sockpath := viper.GetString("sockpath")
		_, err := os.Stat(sockpath)
		if err != nil {
			log.Fatal(err)
		}

		c, err := net.Dial("unix", sockpath)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()

		for {
			text, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Println(err)
				break
			}
			c.Write([]byte(text))

			text, _ = bufio.NewReader(c).ReadString('\n')
			fmt.Print(text)

			if strings.TrimSpace(text) == "STOP" {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
