package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	CommandStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).MarginTop(1)
	OutputStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginBottom(1)
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)
