package main

import (
	"github.com/sarkarshuvojit/bubblebook/examples/basic/models"
	"github.com/sarkarshuvojit/bubblebook/pkg/bubblebook"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Root-level component (ungrouped)
	bubblebook.Register("Welcome", func() tea.Model {
		return models.NewToggleModel()
	})

	// Interactive components group
	bubblebook.Group("Interactive", func() {
		bubblebook.Register("Counter", func() tea.Model {
			return models.NewCounterModel()
		})

		bubblebook.Register("Toggle", func() tea.Model {
			return models.NewToggleModel()
		})
	})

	// Feedback components group
	bubblebook.Group("Feedback", func() {
		bubblebook.Register("Spinner", func() tea.Model {
			return models.NewSpinnerModel()
		})
	})

	bubblebook.Start()
}
