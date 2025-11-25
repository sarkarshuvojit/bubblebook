package main

import (
	"github.com/sarkarshuvojit/bubblebook/examples/basic/models"
	"github.com/sarkarshuvojit/bubblebook/pkg/bubblebook"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	bubblebook.Register("Spinner", func() tea.Model {
		return models.NewSpinnerModel()
	})

	bubblebook.Register("Counter", func() tea.Model {
		return models.NewCounterModel()
	})

	bubblebook.Register("Toggle", func() tea.Model {
		return models.NewToggleModel()
	})

	bubblebook.Start()
}
