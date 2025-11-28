package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/progress"
)

type ProgressModel struct {
	progress progress.Model
	percent  float64
}

func NewProgressModel() ProgressModel {
	return ProgressModel{
		progress: progress.New(progress.WithDefaultGradient()),
		percent:  0.0,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return nil
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "+", "right", "l":
			m.percent += 0.1
			if m.percent > 1.0 {
				m.percent = 1.0
			}
		case "-", "left", "h":
			m.percent -= 0.1
			if m.percent < 0.0 {
				m.percent = 0.0
			}
		case "r":
			m.percent = 0.0
		}
	}
	return m, nil
}

func (m ProgressModel) View() string {
	pad := strings.Repeat(" ", 2)
	return "\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n" +
		pad + fmt.Sprintf("%.0f%%", m.percent*100) + "\n\n" +
		pad + "(Press '+' or 'l' to increase, '-' or 'h' to decrease, 'r' to reset)\n\n"
}
