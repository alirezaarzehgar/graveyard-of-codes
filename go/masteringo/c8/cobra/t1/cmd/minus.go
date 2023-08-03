/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// minusCmd represents the minus command
var minusCmd = &cobra.Command{
	Use:     "minus",
	Aliases: []string{"min", "^"},
	Short:   "x minus y",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("minus called")

		x, _ := cmd.Flags().GetInt("x")
		y, _ := cmd.Flags().GetInt("y")
		fmt.Printf("%d - %d = %d\n", x, y, x-y)
	},
}

func init() {
	rootCmd.AddCommand(minusCmd)

	minusCmd.PersistentFlags().Int("x", 0, "get x")
	minusCmd.MarkPersistentFlagRequired("x")
	minusCmd.PersistentFlags().Int("y", 0, "get y")
	minusCmd.MarkPersistentFlagRequired("y")
}
