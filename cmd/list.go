package cmd

import (
	"fmt"

	"github.com/JulianH99/bills/internal/data"
	"github.com/JulianH99/bills/internal/services"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all bills",
	RunE: func(cmd *cobra.Command, args []string) error {
		db := data.NewDatabase("bills.db")
		err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		bills, err := services.ListBills(db.Instance())
		if err != nil {
			return err
		}

		for _, b := range bills {
			fmt.Println(b.Name)
		}

		return nil
	},
}
