/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:     "div",
	Aliases: []string{"divide", "/", "d"},
	Short:   "divide y to x",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("div called")

		x, _ := cmd.Flags().GetInt("x")
		y, _ := cmd.Flags().GetInt("y")
		fmt.Printf("%d / %d = %d\n", x, y, y/x)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)

	divCmd.PersistentFlags().Int("x", 0, "get x")
	divCmd.MarkPersistentFlagRequired("x")
	divCmd.PersistentFlags().Int("y", 0, "get y")
	divCmd.MarkPersistentFlagRequired("y")
}
