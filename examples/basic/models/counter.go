package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// counterModel is a simple counter component.
type counterModel struct {
	count int
}

func NewCounterModel() counterModel {
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
