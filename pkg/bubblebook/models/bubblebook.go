package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Pane represents which pane is currently focused
type Pane int

const (
	PaneList Pane = iota
	PanePreview
)

// ComponentEntry represents a registered component
type ComponentEntry struct {
	Name    string
	Factory func() tea.Model
}

// BubblebookModel is the main application model
type BubblebookModel struct {
	// Layout dimensions
	width  int
	height int

	// Configuration
	sidebarWidth int

	// State
	entries       []Entry
	selectedIndex int
	focusedPane   Pane
	showHelp      bool

	// Sub-models
	componentList *ComponentListModel
	preview       *PreviewModel
}

// NewBubblebookModel creates a new instance of the main application model
func NewBubblebookModel(entries []Entry) BubblebookModel {
	return BubblebookModel{
		sidebarWidth:  30,
		entries:       entries,
		selectedIndex: 0,
		focusedPane:   PaneList,
		componentList: NewComponentListModel(entries),
		preview:       NewPreviewModel(),
	}
}

// Init initializes the model
func (m BubblebookModel) Init() tea.Cmd {
	// Load the first component if available
	if CountVisibleEntries(m.entries) > 0 {
		return m.loadComponent(0)
	}
	return nil
}

// Update handles messages
func (m BubblebookModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Update component list size
		m.componentList.SetSize(m.sidebarWidth-2, m.height-2)

		// Update preview size
		previewWidth := m.width - m.sidebarWidth
		m.preview.SetSize(previewWidth-2, m.height-2)

		// Forward resize to active component if any
		if m.preview.HasComponent() {
			cmd = m.preview.ForwardMessage(msg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			// Don't quit if help is showing, just close help
			if m.showHelp {
				m.showHelp = false
				return m, nil
			}
			return m, tea.Quit

		case "?":
			// Toggle help screen
			m.showHelp = !m.showHelp
			return m, nil

		case "tab":
			// Don't switch focus if help is showing
			if m.showHelp {
				return m, nil
			}
			// Toggle focus between panes
			if m.focusedPane == PaneList {
				m.focusedPane = PanePreview
			} else {
				m.focusedPane = PaneList
			}
			m.componentList.SetFocused(m.focusedPane == PaneList)
			m.preview.SetFocused(m.focusedPane == PanePreview)

		case "esc":
			// If help is showing, close it
			if m.showHelp {
				m.showHelp = false
				return m, nil
			}
			// Always return to list
			m.focusedPane = PaneList
			m.componentList.SetFocused(true)
			m.preview.SetFocused(false)

		default:
			// Don't route messages if help is showing
			if m.showHelp {
				return m, nil
			}
			// Route message based on focused pane
			if m.focusedPane == PaneList {
				var listCmd tea.Cmd
				*m.componentList, listCmd = m.componentList.Update(msg)
				if listCmd != nil {
					cmds = append(cmds, listCmd)
				}

				// Check if selection changed
				if m.componentList.SelectedIndex() != m.selectedIndex {
					m.selectedIndex = m.componentList.SelectedIndex()
					cmd = m.loadComponent(m.selectedIndex)
					if cmd != nil {
						cmds = append(cmds, cmd)
					}
				}
			} else if m.focusedPane == PanePreview {
				// Forward to preview (which forwards to active component)
				cmd = m.preview.ForwardMessage(msg)
				if cmd != nil {
					cmds = append(cmds, cmd)
				}
			}
		}

	default:
		// Forward all other messages to preview if it has a component
		if m.preview.HasComponent() {
			cmd = m.preview.ForwardMessage(msg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}

	return m, tea.Batch(cmds...)
}

// View renders the application
func (m BubblebookModel) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	// If help is showing, render help instead
	if m.showHelp {
		return RenderHelp(m.width, m.height)
	}

	// Render component list
	listView := m.componentList.View()

	// Render preview
	previewView := m.preview.View()

	// Join horizontally
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		listView,
		previewView,
	)
}

// loadComponent loads a component by index
func (m BubblebookModel) loadComponent(index int) tea.Cmd {
	factory, name, ok := FindComponentByIndex(m.entries, index)
	if !ok {
		return nil
	}

	component := factory()

	// Set the component in the preview
	return m.preview.LoadComponent(component, name)
}
