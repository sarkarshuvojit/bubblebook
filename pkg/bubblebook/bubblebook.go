package bubblebook

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sarkarshuvojit/bubblebook/pkg/bubblebook/models"
)

// ComponentFactory is a function that returns a new instance of a Bubble Tea model.
type ComponentFactory func() tea.Model

var components []models.ComponentEntry

// Register adds a component to the Bubblebook registry.
func Register(name string, factory ComponentFactory) {
	components = append(components, models.ComponentEntry{
		Name:    name,
		Factory: factory,
	})
}

// Start launches the Bubblebook TUI.
func Start() {
	// Create the main model
	model := models.NewBubblebookModel(components)

	// Create the program
	program := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// Run the program
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running bubblebook: %v\n", err)
		os.Exit(1)
	}
}
