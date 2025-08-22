package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rivo/tview"
)

// model is a minimal Bubble Tea model that simply renders
// a static loading message.
type model struct{}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m model) View() string { return "Loading..." }

func main() {
	app := tview.NewApplication()

	// Instantiate the Bubble Tea model and obtain its view.
	loadingView := model{}.View()

	// Create a grid with 2 rows and 2 columns.
	// The left column and bottom row have fixed sizes similar to the diagram.
	grid := tview.NewGrid().
		SetRows(0, 3).     // Bottom bar height 3
		SetColumns(20, 0). // Left panel width 20
		SetBorders(true)   // Draw borders between cells

	left := tview.NewBox()
	// Display the Bubble Tea model's view in a TextView on the right side.
	mainArea := tview.NewTextView().SetText(loadingView)
	bottom := tview.NewBox()

	grid.AddItem(left, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(mainArea, 0, 1, 1, 1, 0, 0, true)
	grid.AddItem(bottom, 1, 0, 1, 2, 0, 0, false)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
