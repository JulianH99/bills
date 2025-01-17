package ui

import "github.com/charmbracelet/lipgloss"

var messageStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	Padding(1)

func Message(text string, isError bool) string {
	var style lipgloss.Style
	if isError {
		style = messageStyle.BorderForeground(lipgloss.Color("#ff0000"))
	} else {
		style = messageStyle.BorderForeground(lipgloss.Color("#ffffff"))
	}

	return style.Render(text)
}
