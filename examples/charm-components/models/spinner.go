package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
)

type SpinnerModel struct {
	spinner spinner.Model
}

func NewSpinnerModel() SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return SpinnerModel{spinner: s}
}

func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m SpinnerModel) View() string {
	return fmt.Sprintf("\n\n   %s Loading...\n\n", m.spinner.View())
}
