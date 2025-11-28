package models

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	helpTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Padding(0, 0, 1, 0)

	helpSectionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("141")).
				Bold(true).
				Padding(1, 0, 0, 0)

	helpKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212")).
			Bold(true)

	helpDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("246"))

	helpBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			Padding(2, 4).
			Margin(1, 2)
)

// RenderHelp renders the help screen
func RenderHelp(width, height int) string {
	var b strings.Builder

	// Title
	b.WriteString(helpTitleStyle.Render("Bubblebook - Keyboard Shortcuts"))
	b.WriteString("\n")

	// Navigation section
	b.WriteString(helpSectionStyle.Render("Navigation"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  ↑/k, ↓/j  "))
	b.WriteString(helpDescStyle.Render("Navigate component list"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  g, G      "))
	b.WriteString(helpDescStyle.Render("Jump to top/bottom"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  enter     "))
	b.WriteString(helpDescStyle.Render("Select component"))
	b.WriteString("\n")

	// Folder navigation section
	b.WriteString(helpSectionStyle.Render("Folders"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  →/l/enter "))
	b.WriteString(helpDescStyle.Render("Expand folder"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  ←/h       "))
	b.WriteString(helpDescStyle.Render("Collapse folder / Jump to parent"))
	b.WriteString("\n")

	// Focus section
	b.WriteString(helpSectionStyle.Render("Focus"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  tab       "))
	b.WriteString(helpDescStyle.Render("Switch between list and preview"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  esc       "))
	b.WriteString(helpDescStyle.Render("Return to component list"))
	b.WriteString("\n")

	// General section
	b.WriteString(helpSectionStyle.Render("General"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  ?         "))
	b.WriteString(helpDescStyle.Render("Toggle this help screen"))
	b.WriteString("\n")
	b.WriteString(helpKeyStyle.Render("  q, ctrl+c "))
	b.WriteString(helpDescStyle.Render("Quit"))
	b.WriteString("\n")

	// Component interaction
	b.WriteString(helpSectionStyle.Render("Component Interaction"))
	b.WriteString("\n")
	b.WriteString(helpDescStyle.Render("  When preview is focused, all keys are forwarded to the active component."))
	b.WriteString("\n")
	b.WriteString(helpDescStyle.Render("  Each component has its own keyboard shortcuts - check the component"))
	b.WriteString("\n")
	b.WriteString(helpDescStyle.Render("  documentation for details."))
	b.WriteString("\n")

	content := b.String()

	return helpBoxStyle.
		Width(width - 8).
		Height(height - 4).
		Render(content)
}
