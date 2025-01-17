package cmd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/JulianH99/bills/internal/services"
	"github.com/JulianH99/bills/internal/ui"
	"github.com/spf13/cobra"
)

func addCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new bill",
		Long: `Add a new bill providing name and day of month to be paid recurringly.
	Example:
		bills add Electricity 12 # Adds a bill named "Electricity" to be paid on the 12th of each month
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check that the args are valid
			if len(args) != 2 {
				return fmt.Errorf("Expected name and day of month of the bill. See usage for examples")
			}

			var (
				name       = args[0]
				dayOfMonth = args[1]
			)

			domInt, err := strconv.Atoi(dayOfMonth)
			if err != nil {
				return fmt.Errorf("Plase provide a valid number for the day of month")
			}

			if domInt < 1 || domInt > 31 {
				return fmt.Errorf("Day of month must be between 1 and 31")
			}

			// TODO: create tables only on first run
			err = services.SaveBill(db, services.Bill{Name: name, DayOfMonth: domInt, Paid: false})
			if err != nil {
				return err
			}

			fmt.Print(ui.Message(fmt.Sprintf("Bill %s added", name), false))

			return nil
		},
	}
}
