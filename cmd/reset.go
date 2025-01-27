package cmd

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JulianH99/bills/internal/services"
	"github.com/JulianH99/bills/internal/ui"
	"github.com/spf13/cobra"
)

func resetCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use: "reset",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentDay := time.Now().Day()

			fmt.Println(currentDay)

			if currentDay == 1 {
				err := services.Reset(db)
				if err != nil {
					fmt.Println(ui.Message("Error reseting bills", true))
				}
			}
			return nil
		},
	}
}
