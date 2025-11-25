package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

// toggleModel is a simple checkbox toggle component.
type toggleModel struct {
	checked bool
}

func NewToggleModel() toggleModel {
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
		checkbox = "[x]"
	}
	return "Toggle: " + checkbox + "\n\n(Press space or enter to toggle)"
}
