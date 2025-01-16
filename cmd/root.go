package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bills",
	Short: "A CLI for managing recurring bills",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
}

func ExecuteRoodCommand() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("An error ocurred. Error: %v", err)
		os.Exit(1)
	}
}
