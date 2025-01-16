package cmd

import (
	"database/sql"
	"fmt"

	"github.com/JulianH99/bills/internal/services"
	"github.com/spf13/cobra"
)

func removeCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "remove",
		Short: "Removes a bill by name",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Only the name of the bill to be removed should be provided")
			}

			name := args[0]

			rows, err := services.RemoveBill(db, name)
			if err != nil {
				return err
			}

			if rows == 0 {
				fmt.Printf("Bill \"%s\" not found\n", name)
				return nil
			}

			fmt.Printf("Bill \"%s\" removed\n", name)
			return nil
		},
	}
}
