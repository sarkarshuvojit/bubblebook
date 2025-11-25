package main

import (
	"fmt"
	"time"

	"github.com/sarkarshuvojit/bubblebook/bubblebook"

	tea "github.com/charmbracelet/bubbletea"
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

func main() {
	bubblebook.Register("Spinner", func() tea.Model {
		return newModel()
	})

	bubblebook.Register("Counter", func() tea.Model {
		return newCounterModel()
	})

	bubblebook.Register("Toggle", func() tea.Model {
		return newToggleModel()
	})

	bubblebook.Start()
}

// --- Counter Component ---

type counterModel struct {
	count int
}

func newCounterModel() counterModel {
	return counterModel{}
}

func (m counterModel) Init() tea.Cmd {
	return nil
}

func (m counterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "+", "k", "up":
			m.count++
		case "-", "j", "down":
			m.count--
		}
	}
	return m, nil
}

func (m counterModel) View() string {
	return "Count: " + fmt.Sprintf("%d", m.count) + "\n\n(Press '+' or 'k' to increment, '-' or 'j' to decrement)"
}

// --- Toggle Component ---

type toggleModel struct {
	checked bool
}

func newToggleModel() toggleModel {
	return toggleModel{}
}

func (m toggleModel) Init() tea.Cmd {
	return nil
}

func (m toggleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ", "enter":
			m.checked = !m.checked
		}
	}
	return m, nil
}

func (m toggleModel) View() string {
	checkbox := "[ ]"
	if m.checked {
		checkbox = "[*]"
	}
	return "Toggle: " + checkbox + "\n\n(Press space or enter to toggle)"
}

