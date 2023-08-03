/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// plusCmd represents the plus command
var plusCmd = &cobra.Command{
	Use:     "plus",
	Aliases: []string{"p", "add", "+"},
	Short:   "add x to y",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plus called")
		x, _ := cmd.Flags().GetInt("x")
		y, _ := cmd.Flags().GetInt("y")

		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	},
}

func init() {
	rootCmd.AddCommand(plusCmd)

	plusCmd.PersistentFlags().Int("x", 0, "get x")
	plusCmd.MarkPersistentFlagRequired("x")
	plusCmd.PersistentFlags().Int("y", 0, "get y")
	plusCmd.MarkPersistentFlagRequired("y")
}
