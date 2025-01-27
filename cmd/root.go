package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bills",
	Short: "A CLI for managing recurring bills",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

func CreateRootCommand(db *sql.DB) *cobra.Command {
	rootCmd.AddCommand(addCmd(db))
	rootCmd.AddCommand(listCmd(db))
	rootCmd.AddCommand(removeCmd(db))
	rootCmd.AddCommand(paidCmd(db))
	rootCmd.AddCommand(setCmd(db))
	rootCmd.AddCommand(notifyCmd(db))
	rootCmd.AddCommand(resetCmd(db))
	return rootCmd
}
