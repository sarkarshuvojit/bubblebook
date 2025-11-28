package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type TextInputModel struct {
	textInput textinput.Model
}

func NewTextInputModel() TextInputModel {
	ti := textinput.New()
	ti.Placeholder = "Type something..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	return TextInputModel{textInput: ti}
}

func (m TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m TextInputModel) View() string {
	return fmt.Sprintf(
		"\n\n   %s\n\n   %s\n\n",
		m.textInput.View(),
		"(Type to enter text)",
	)
}
