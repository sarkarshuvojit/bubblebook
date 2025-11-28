package main

import (
	"github.com/sarkarshuvojit/bubblebook/examples/charm-components/models"
	"github.com/sarkarshuvojit/bubblebook/pkg/bubblebook"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	bubblebook.Register("Spinner", func() tea.Model {
		return models.NewSpinnerModel()
	})

	bubblebook.Register("Progress Bar", func() tea.Model {
		return models.NewProgressModel()
	})

	bubblebook.Register("Text Input", func() tea.Model {
		return models.NewTextInputModel()
	})

	bubblebook.Register("List", func() tea.Model {
		return models.NewListModel()
	})

	bubblebook.Start()
}
