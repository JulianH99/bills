package cmd

import (
	"database/sql"
	"fmt"

	"github.com/JulianH99/bills/internal/services"
	"github.com/JulianH99/bills/internal/ui"
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
				fmt.Println(ui.Message(fmt.Sprintf("Bill \"%s\" not found", name), false))

				return nil
			}

			fmt.Println(ui.Message(fmt.Sprintf("Paid bill \"%s\"", name), false))

			return nil
		},
	}
}
