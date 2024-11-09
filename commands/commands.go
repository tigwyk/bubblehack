package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Define a map to hold command functions
var CommandMap = map[string]func() *cobra.Command{
	"echo": EchoCommand,
	"quit": QuitCommand,
	// Add other commands here
}

func EchoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "echo",
		Short: "Display a line of text",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.Join(args, " "))
		},
	}
}

func QuitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "quit",
		Short: "Exit the game",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Exiting bubblehack...")
			os.Exit(0)
		},
	}
}
