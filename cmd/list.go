package cmd

import (
	"database/sql"
	"fmt"

	"github.com/JulianH99/bills/internal/services"
	"github.com/JulianH99/bills/internal/ui"
	"github.com/spf13/cobra"
)

func listCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all bills",
		RunE: func(cmd *cobra.Command, args []string) error {
			bills, err := services.ListBills(db)
			if err != nil {
				return err
			}

			table := ui.PrintAsTable(bills)

			fmt.Printf("%s\n", table)

			return nil
		},
	}
}
