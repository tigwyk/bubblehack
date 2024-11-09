package main

import (
	"bubblehack/commands"
	"bubblehack/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "bubblehack",
	Short:        "A terminal-based hacking game",
	Long:         "bubblehack is a terminal-based hacking game written in Go.",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(models.InitialModel(), tea.WithAltScreen())
		if err := p.Start(); err != nil {
			panic(err)
		}
	},
}

func main() {
	// Dynamically add commands from the map
	for _, cmdFunc := range commands.CommandMap {
		rootCmd.AddCommand(cmdFunc())
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
