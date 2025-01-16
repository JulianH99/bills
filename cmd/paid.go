package cmd

import (
	"database/sql"
	"fmt"

	"github.com/JulianH99/bills/internal/services"
	"github.com/spf13/cobra"
)

func paidCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:     "paid",
		Short:   "Marks a bill as paid (for the current month. At the start of the month everything is reset)",
		Example: "bills paid electricity",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Expected the name of the bill to be paid. See usage for examples")
			}

			name := args[0]

			rows, err := services.Paid(db, name)
			if err != nil {
				return err
			}

			if rows == 0 {
				fmt.Printf("Bill \"%s\" not found\n", name)
				return nil
			}

			fmt.Printf("Paid bill \"%s\"\n", name)

			return nil
		},
	}
}
