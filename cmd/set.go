package cmd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/JulianH99/bills/internal/services"
	"github.com/JulianH99/bills/internal/ui"
	"github.com/spf13/cobra"
)

func setCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Set the date for a given bill",
		RunE: func(cmd *cobra.Command, args []string) error {
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

			bill := services.Bill{Name: name, DayOfMonth: domInt}
			err = services.Update(db, bill)
			if err != nil {
				return err
			}

			cardinal := "th"
			switch domInt {
			case 1, 21, 31:
				cardinal = "st"
			case 2, 22:
				cardinal = "nd"
			case 3, 23:
				cardinal = "rd"
			}

			fmt.Println(ui.Message(fmt.Sprintf("Bill \"%s\" updated to %s%s of each month", name, dayOfMonth, cardinal), false))

			return nil
		},
	}
}
