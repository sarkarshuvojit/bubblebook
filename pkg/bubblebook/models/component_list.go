package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	listBorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1)

	listBorderFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("205")).
				Padding(0, 1)

	listTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("205")).
				Bold(true).
				PaddingLeft(2)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("246")).
			PaddingLeft(2)

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))
)

// ComponentListModel handles the component list sidebar
type ComponentListModel struct {
	components    []ComponentEntry
	selectedIndex int
	width         int
	height        int
	focused       bool
	scrollOffset  int
}

// NewComponentListModel creates a new component list model
func NewComponentListModel(components []ComponentEntry) *ComponentListModel {
	return &ComponentListModel{
		components:    components,
		selectedIndex: 0,
		focused:       true,
		scrollOffset:  0,
	}
}

// SetSize updates the dimensions
func (m *ComponentListModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// SetFocused sets the focus state
func (m *ComponentListModel) SetFocused(focused bool) {
	m.focused = focused
}

// SelectedIndex returns the current selected index
func (m *ComponentListModel) SelectedIndex() int {
	return m.selectedIndex
}

// Update handles messages
func (m ComponentListModel) Update(msg tea.Msg) (ComponentListModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selectedIndex > 0 {
				m.selectedIndex--
				m.ensureVisible()
			}
		case "down", "j":
			if m.selectedIndex < len(m.components)-1 {
				m.selectedIndex++
				m.ensureVisible()
			}
		case "g":
			// Go to top
			m.selectedIndex = 0
			m.scrollOffset = 0
		case "G":
			// Go to bottom
			m.selectedIndex = len(m.components) - 1
			m.ensureVisible()
		case "home":
			m.selectedIndex = 0
			m.scrollOffset = 0
		case "end":
			m.selectedIndex = len(m.components) - 1
			m.ensureVisible()
		}
	}

	return m, nil
}

// ensureVisible adjusts scroll offset to keep selected item visible
func (m *ComponentListModel) ensureVisible() {
	visibleLines := m.height - 4 // Account for border and title

	if m.selectedIndex < m.scrollOffset {
		m.scrollOffset = m.selectedIndex
	} else if m.selectedIndex >= m.scrollOffset+visibleLines {
		m.scrollOffset = m.selectedIndex - visibleLines + 1
	}
}

// View renders the component list
func (m ComponentListModel) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	var b strings.Builder

	// Title
	title := listTitleStyle.Render("Components")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Calculate visible range
	visibleLines := m.height - 4 // Account for border, title, and padding
	if visibleLines < 1 {
		visibleLines = 1
	}

	startIdx := m.scrollOffset
	endIdx := m.scrollOffset + visibleLines
	if endIdx > len(m.components) {
		endIdx = len(m.components)
	}

	// Render items
	for i := startIdx; i < endIdx; i++ {
		component := m.components[i]
		cursor := "  "
		style := normalItemStyle

		if i == m.selectedIndex {
			cursor = cursorStyle.Render("▶ ")
			style = selectedItemStyle
		}

		line := cursor + style.Render(component.Name)
		b.WriteString(line)
		b.WriteString("\n")
	}

	// Add scroll indicators if needed
	if m.scrollOffset > 0 {
		// Can scroll up
		b.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("  ▲ more"))
	}
	if endIdx < len(m.components) {
		// Can scroll down
		b.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("  ▼ more"))
	}

	content := b.String()

	// Apply border
	borderStyle := listBorderStyle
	if m.focused {
		borderStyle = listBorderFocusedStyle
	}

	return borderStyle.
		Width(m.width).
		Height(m.height).
		Render(content)
}
