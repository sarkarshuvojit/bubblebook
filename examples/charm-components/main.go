package main

import (
	"github.com/sarkarshuvojit/bubblebook/examples/charm-components/models"
	"github.com/sarkarshuvojit/bubblebook/pkg/bubblebook"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Input components group
	bubblebook.Group("Input", func() {
		bubblebook.Register("Text Input", func() tea.Model {
			return models.NewTextInputModel()
		})
	})

	// Data display group
	bubblebook.Group("Data Display", func() {
		bubblebook.Register("List", func() tea.Model {
			return models.NewListModel()
		})
	})

	// Feedback group with nested groups
	bubblebook.Group("Feedback", func() {
		bubblebook.Group("Progress", func() {
			bubblebook.Register("Progress Bar", func() tea.Model {
				return models.NewProgressModel()
			})
		})

		bubblebook.Group("Loading", func() {
			bubblebook.Register("Spinner", func() tea.Model {
				return models.NewSpinnerModel()
			})
		})
	})

	bubblebook.Start()
}
