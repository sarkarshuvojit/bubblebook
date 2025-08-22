package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create a grid with 2 rows and 2 columns.
	// The left column and bottom row have fixed sizes similar to the diagram.
	grid := tview.NewGrid().
		SetRows(0, 3).     // Bottom bar height 3
		SetColumns(20, 0). // Left panel width 20
		SetBorders(true)   // Draw borders between cells

	left := tview.NewBox()
	mainArea := tview.NewBox()
	bottom := tview.NewBox()

	grid.AddItem(left, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(mainArea, 0, 1, 1, 1, 0, 0, true)
	grid.AddItem(bottom, 1, 0, 1, 2, 0, 0, false)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
