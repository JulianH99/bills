package ui

import (
	"strconv"

	"github.com/JulianH99/bills/internal/services"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func PrintAsTable(bills []services.Bill) string {
	columns := []table.Column{
		{Title: "Name", Width: 15},
		{Title: "Day of month", Width: 15},
		{Title: "Paid", Width: 5},
	}

	rows := make([]table.Row, len(bills))
	for i, bill := range bills {
		var paid string
		if bill.Paid {
			paid = "Yes"
		} else {
			paid = "No"
		}

		rows[i] = table.Row{bill.Name, strconv.Itoa(bill.DayOfMonth), paid}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(len(rows)+2+len(rows)/2),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true)

	s.Cell = s.Cell.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true)

	t.SetStyles(s)

	return t.View()
}
