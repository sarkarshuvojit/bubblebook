package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rivo/tview"
)

// model shows a simple ASCII spinner before rendering a loading message.
type model struct {
	frames []string
	frame  int
	done   bool
}

func newModel() model {
	return model{frames: []string{"|", "/", "-", "\\"}}
}

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(time.Time) tea.Msg { return tickMsg{} })
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		if m.frame >= len(m.frames)*8 {
			m.done = true
			return m, tea.Quit
		}
		m.frame++
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	if m.done {
		return "Loading..."
	}
	return m.frames[m.frame%len(m.frames)]
}

type tviewWriter struct {
	app *tview.Application
	tv  *tview.TextView
}

func (w tviewWriter) Write(p []byte) (int, error) {
	s := string(p)
	w.app.QueueUpdateDraw(func() {
		w.tv.SetText(s)
	})
	return len(p), nil
}

func main() {
	app := tview.NewApplication()

	// Create a grid with 2 rows and 2 columns.
	// The left column and bottom row have fixed sizes similar to the diagram.
	grid := tview.NewGrid().
		SetRows(0, 3).     // Bottom bar height 3
		SetColumns(20, 0). // Left panel width 20
		SetBorders(true)   // Draw borders between cells

	left := tview.NewBox()
	mainArea := tview.NewTextView()
	bottom := tview.NewBox()

	grid.AddItem(left, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(mainArea, 0, 1, 1, 1, 0, 0, true)
	grid.AddItem(bottom, 1, 0, 1, 2, 0, 0, false)

	go func() {
		w := tviewWriter{app: app, tv: mainArea}
		if _, err := tea.NewProgram(newModel(), tea.WithOutput(w)).Run(); err != nil {
			panic(err)
		}
	}()

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
