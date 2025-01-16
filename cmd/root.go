package cmd

import (
	"database/sql"
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

func ExecuteRootCommand(db *sql.DB) {
	rootCmd.AddCommand(addCmd(db))
	rootCmd.AddCommand(listCmd(db))
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("An error ocurred. Error: %v", err)
		os.Exit(1)
	}
}
