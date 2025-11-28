package bubblebook

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sarkarshuvojit/bubblebook/pkg/bubblebook/models"
)

// ComponentFactory is a function that returns a new instance of a Bubble Tea model.
type ComponentFactory func() tea.Model

var entries []models.Entry
var groupStack []*models.Entry

// Register adds a component to the current group (or root if no group).
func Register(name string, factory ComponentFactory) {
	entry := models.Entry{
		Type:    models.EntryTypeComponent,
		Name:    name,
		Factory: factory,
	}

	if len(groupStack) > 0 {
		// Add to current group
		currentGroup := groupStack[len(groupStack)-1]
		currentGroup.Children = append(currentGroup.Children, entry)
	} else {
		// Add to root
		entries = append(entries, entry)
	}
}

// Group creates a new group and executes the register function within it.
func Group(name string, register func()) {
	group := &models.Entry{
		Type:     models.EntryTypeGroup,
		Name:     name,
		Children: []models.Entry{},
		Expanded: true, // Groups expanded by default
	}

	// Push group onto stack
	groupStack = append(groupStack, group)

	// Execute registration function
	register()

	// Pop group from stack
	groupStack = groupStack[:len(groupStack)-1]

	// Add group to parent (or root)
	if len(groupStack) > 0 {
		parent := groupStack[len(groupStack)-1]
		parent.Children = append(parent.Children, *group)
	} else {
		entries = append(entries, *group)
	}
}

// Start launches the Bubblebook TUI.
func Start() {
	// Create the main model
	model := models.NewBubblebookModel(entries)

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
