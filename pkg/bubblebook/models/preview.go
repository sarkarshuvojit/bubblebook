package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	previewBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2)

	previewBorderFocusedStyle = lipgloss.NewStyle().
					Border(lipgloss.RoundedBorder()).
					BorderForeground(lipgloss.Color("205")).
					Padding(1, 2)

	previewTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("205")).
				Bold(true)

	emptyStateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)
)

// PreviewModel handles the component preview area
type PreviewModel struct {
	component     tea.Model
	componentName string
	hasComponent  bool
	width         int
	height        int
	focused       bool
}

// NewPreviewModel creates a new preview model
func NewPreviewModel() *PreviewModel {
	return &PreviewModel{
		hasComponent: false,
		focused:      false,
	}
}

// SetSize updates the dimensions
func (m *PreviewModel) SetSize(width, height int) {
	m.width = width
	m.height = height

	// Forward resize to component if active
	if m.hasComponent && m.component != nil {
		msg := tea.WindowSizeMsg{
			Width:  width - 4, // Account for padding
			Height: height - 4,
		}
		var cmd tea.Cmd
		m.component, cmd = m.component.Update(msg)
		_ = cmd // We'll handle this in the main update loop
	}
}

// SetFocused sets the focus state
func (m *PreviewModel) SetFocused(focused bool) {
	m.focused = focused
}

// HasComponent returns whether a component is loaded
func (m *PreviewModel) HasComponent() bool {
	return m.hasComponent
}

// LoadComponent loads a new component into the preview
func (m *PreviewModel) LoadComponent(component tea.Model, name string) tea.Cmd {
	m.component = component
	m.componentName = name
	m.hasComponent = true

	// Initialize the component
	if m.component != nil {
		// Send initial window size
		msg := tea.WindowSizeMsg{
			Width:  m.width - 4,
			Height: m.height - 4,
		}
		var cmd tea.Cmd
		m.component, cmd = m.component.Update(msg)

		// Return the component's Init command
		return tea.Batch(cmd, m.component.Init())
	}

	return nil
}

// ForwardMessage forwards a message to the active component
func (m *PreviewModel) ForwardMessage(msg tea.Msg) tea.Cmd {
	if !m.hasComponent || m.component == nil {
		return nil
	}

	var cmd tea.Cmd
	m.component, cmd = m.component.Update(msg)
	return cmd
}

// View renders the preview area
func (m PreviewModel) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	var content string

	if m.hasComponent && m.component != nil {
		// Render the active component
		componentView := m.component.View()

		// Add title
		title := previewTitleStyle.Render(m.componentName)

		// Add help text if focused
		var help string
		if m.focused {
			help = helpStyle.Render("Press ESC to return to list • Press ? for help • Press q to quit")
		} else {
			help = helpStyle.Render("Press TAB to focus preview • Press ? for help")
		}

		// Combine title, component view, and help
		var b strings.Builder
		b.WriteString(title)
		b.WriteString("\n\n")
		b.WriteString(componentView)
		b.WriteString("\n\n")
		b.WriteString(help)

		content = b.String()
	} else {
		// Empty state
		content = emptyStateStyle.Render("No component selected\n\nSelect a component from the list to preview it here.")
	}

	// Apply border
	borderStyle := previewBorderStyle
	if m.focused {
		borderStyle = previewBorderFocusedStyle
	}

	return borderStyle.
		Width(m.width).
		Height(m.height).
		Render(content)
}
