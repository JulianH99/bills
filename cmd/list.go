package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/JulianH99/bills/internal/services"
	"github.com/JulianH99/bills/internal/ui"
	"github.com/spf13/cobra"
)

var jsonFlag bool

func listCmd(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all bills",
		RunE: func(cmd *cobra.Command, args []string) error {
			bills, err := services.ListBills(db)
			if err != nil {
				return err
			}

			if jsonFlag {
				jsonContent, err := json.Marshal(bills)
				if err != nil {
					return err
				}

				fmt.Println(string(jsonContent))
			} else {
				table := ui.PrintAsTable(bills)

				fmt.Printf("%s\n", table)
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&jsonFlag, "json", false, "Prints the list as json formatted content")

	return cmd
}
