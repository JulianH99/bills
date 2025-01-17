package ui

import "github.com/charmbracelet/lipgloss"

func Message(text string, isError bool) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1)
	if isError {
		style = style.BorderForeground(lipgloss.Color("#ff0000"))
	} else {
		style = style.BorderForeground(lipgloss.Color("#ffffff"))
	}

	return style.Render(text)
}
