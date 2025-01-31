package cmd

import (
	"database/sql"
	"fmt"
	"os/exec"
	"strings"

	"github.com/JulianH99/bills/internal/config"
	"github.com/JulianH99/bills/internal/services"
	"github.com/spf13/cobra"
)

const (
	notifyCmdName = "notify"
	severity      = "critical"
	appName       = "gbills"
	expireTime    = "3000"
)

func notifyCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use: "notify",
		RunE: func(cmd *cobra.Command, args []string) error {
			daysBefore := 3
			bills, err := services.ListBillsSoonToBePaid(db, daysBefore)
			if err != nil {
				return err
			}

			if len(bills) == 0 {
				return nil
			}

			sb := strings.Builder{}

			sb.WriteString("Bills to be paid in the following 3 days:\n")
			for _, bill := range bills {
				sb.WriteString(fmt.Sprintf("- %s bill due to be paid on %d\n", bill.Name, bill.DayOfMonth))
			}
			cfg, _ := config.GetConfig()

			fmt.Println("icon", cfg.IconPath)

			// TODO: add windows support
			_, err = exec.Command(notifyCmdName,
				"--urgency", severity,
				"--icon", cfg.IconPath,
				"--app-name", appName,
				"--expire-time", expireTime,
				"Bills to be paid", sb.String(),
			).Output()
			if err != nil {
				// TODO: should log to a file
				return err
			}

			return nil
		},
	}
}
