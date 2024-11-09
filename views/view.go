package views

import (
	"bubblehack/models"
	"bubblehack/styles"
)

func View(m models.Model) string {
	outputView := styles.OutputStyle.Render(m.Output)
	inputView := styles.CommandStyle.Render(m.CommandInput.View())

	if m.Processing {
		return outputView + "\n" + m.Spinner.View() + "\n" + inputView
	}
	return outputView + "\n" + inputView
}
