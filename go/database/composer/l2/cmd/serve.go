/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"hiadvanced/api"
	"hiadvanced/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve application",
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load(".env")

		db, err := storage.NewConnection(storage.MysqlConf{
			Username: os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Hostname: os.Getenv("DB_HOSTNAME"),
			DbName:   os.Getenv("MYSQL_DATABASE"),
			Port:     os.Getenv("PORT"),
		})
		if err != nil {
			log.Fatal(err)
		}
		api.Run(db)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
