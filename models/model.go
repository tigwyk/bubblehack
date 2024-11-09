package models

import (
	"bubblehack/commands"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	CommandInput textinput.Model
	Output       string
	Spinner      spinner.Model
	Processing   bool
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter command"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	sp := spinner.New()
	sp.Spinner = spinner.Dot

	return Model{
		CommandInput: ti,
		Output:       "Welcome to bubblehack!",
		Spinner:      sp,
		Processing:   false,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.Spinner.Tick)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Handle command execution
			args := strings.Fields(m.CommandInput.Value())
			if len(args) > 0 {
				cmdFunc, exists := commands.CommandMap[args[0]]
				if exists {
					m.Processing = true
					cmd := cmdFunc()
					cmd.SetArgs(args[1:])
					cmds = append(cmds, func() tea.Msg {
						err := cmd.Execute()
						if err != nil {
							return err
						}
						m.Processing = false
						return nil
					})
					m.Output = "Command executed: " + args[0]
				} else {
					m.Output = "Unknown command: " + args[0]
				}
			} else {
				m.Output = "No command provided"
			}
			m.CommandInput.SetValue("")
		}

	case spinner.TickMsg:
		if m.Processing {
			var cmd tea.Cmd
			m.Spinner, cmd = m.Spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	if !m.Processing {
		var cmd tea.Cmd
		m.CommandInput, cmd = m.CommandInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	outputView := lipgloss.NewStyle().MarginBottom(1).Render(m.Output)
	inputView := lipgloss.NewStyle().MarginTop(1).Render(m.CommandInput.View())

	if m.Processing {
		return outputView + "\n" + m.Spinner.View() + "\n" + inputView
	}
	return outputView + "\n" + inputView
}
