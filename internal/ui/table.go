package ui

import (
	"strconv"

	"github.com/JulianH99/bills/internal/services"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func PrintAsTable(bills []services.Bill) string {
	paidCellStyle := lipgloss.NewStyle().
		Bold(true)
	unpaidCellStyle := lipgloss.NewStyle().Bold(true)
	fadedCellStyle := lipgloss.NewStyle().Faint(true)
	columns := []table.Column{
		{Title: "Name", Width: 19},
		{Title: "Day of month", Width: 15},
		{Title: "Paid", Width: 15},
	}

	rows := make([]table.Row, len(bills))
	for i, bill := range bills {
		var paid string
		if bill.Paid {
			paid = paidCellStyle.Render("Yes")
		} else {
			paid = unpaidCellStyle.Render("No")
		}
		rows[i] = table.Row{bill.Name, strconv.Itoa(bill.DayOfMonth), paid}
		if bill.Paid {
			for j, cell := range rows[i] {
				rows[i][j] = fadedCellStyle.Render(cell)
			}
		}

	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(len(rows)+4+len(rows)/2),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true)

	s.Selected = s.Selected.Foreground(s.Cell.GetForeground()).
		Bold(false)

	s.Cell = s.Cell.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true)

	t.SetStyles(s)

	return t.View()
}
