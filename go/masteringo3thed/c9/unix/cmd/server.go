/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func echoHandler(c net.Conn) {
	defer c.Close()
	for {
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		c.Write(buf[0:n])
	}
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a UNIX socket server",
	Run: func(cmd *cobra.Command, args []string) {
		sockpath := viper.GetString("sockpath")
		_, err := os.Stat(sockpath)
		if err == nil {
			err = os.Remove(sockpath)
			if err != nil {
				log.Fatal(err)
			}
		}

		l, err := net.Listen("unix", sockpath)
		if err != nil {
			log.Fatal(err)
		}

		for {
			c, err := l.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go echoHandler(c)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
