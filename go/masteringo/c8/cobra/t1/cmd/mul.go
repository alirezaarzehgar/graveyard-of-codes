/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:     "mul",
	Aliases: []string{"multiply", "*"},
	Short:   "multiplty x an y",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mul called")

		x, _ := cmd.Flags().GetInt("x")
		y, _ := cmd.Flags().GetInt("y")
		fmt.Printf("%d * %d = %d\n", x, y, x*y)
	},
}

func init() {
	rootCmd.AddCommand(mulCmd)

	mulCmd.PersistentFlags().Int("x", 0, "get x")
	mulCmd.MarkPersistentFlagRequired("x")
	mulCmd.PersistentFlags().Int("y", 0, "get y")
	mulCmd.MarkPersistentFlagRequired("y")
}
