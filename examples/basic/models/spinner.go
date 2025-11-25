package models

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// spinnerModel shows a simple ASCII spinner before rendering a loading message.
type spinnerModel struct {
	frames []string
	frame  int
	done   bool
}

func NewSpinnerModel() spinnerModel {
	return spinnerModel{frames: []string{"|", "/", "-", "\\"}}
}

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(time.Time) tea.Msg { return tickMsg{} })
}

func (m spinnerModel) Init() tea.Cmd {
	return tick()
}

func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m spinnerModel) View() string {
	if m.done {
		return "Loading..."
	}
	return m.frames[m.frame%len(m.frames)]
}
