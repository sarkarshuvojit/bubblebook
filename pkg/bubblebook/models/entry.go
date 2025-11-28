package models

import tea "github.com/charmbracelet/bubbletea"

// EntryType distinguishes between components and groups
type EntryType int

const (
	EntryTypeComponent EntryType = iota
	EntryTypeGroup
)

// Entry represents either a component or a group
type Entry struct {
	Type     EntryType
	Name     string
	Factory  func() tea.Model // Only used for components
	Children []Entry          // Only used for groups
	Expanded bool             // Only used for groups, tracks UI state
	Level    int              // Depth level for rendering indentation
}

// FlattenEntries converts the tree structure to a flat list for rendering
// Respects the expanded/collapsed state of groups
func FlattenEntries(entries []Entry, level int) []Entry {
	var result []Entry
	for _, entry := range entries {
		entry.Level = level
		result = append(result, entry)

		if entry.Type == EntryTypeGroup && entry.Expanded {
			// Recursively add children if group is expanded
			result = append(result, FlattenEntries(entry.Children, level+1)...)
		}
	}
	return result
}

// FindComponentByIndex finds a component in the flattened list
// Returns the component's factory and name
func FindComponentByIndex(entries []Entry, index int) (func() tea.Model, string, bool) {
	flattened := FlattenEntries(entries, 0)
	if index < 0 || index >= len(flattened) {
		return nil, "", false
	}

	entry := flattened[index]
	if entry.Type == EntryTypeComponent {
		return entry.Factory, entry.Name, true
	}
	return nil, "", false
}

// ToggleGroup toggles the expanded state of a group at the given index in the flattened list
func ToggleGroup(entries []Entry, index int) []Entry {
	flattened := FlattenEntries(entries, 0)
	if index < 0 || index >= len(flattened) {
		return entries
	}

	target := flattened[index]
	if target.Type == EntryTypeGroup {
		// Find and toggle in the original tree
		return toggleInTree(entries, target.Name, target.Level)
	}
	return entries
}

// toggleInTree recursively finds and toggles a group by name and level
func toggleInTree(entries []Entry, name string, targetLevel int) []Entry {
	for i, entry := range entries {
		currentLevel := getEntryLevel(entries[:i+1], 0)
		if entry.Type == EntryTypeGroup && entry.Name == name && currentLevel == targetLevel {
			entries[i].Expanded = !entries[i].Expanded
			return entries
		}
		if entry.Type == EntryTypeGroup && len(entry.Children) > 0 {
			entries[i].Children = toggleInTree(entry.Children, name, targetLevel-1)
		}
	}
	return entries
}

// getEntryLevel calculates the level of an entry in the tree
func getEntryLevel(entries []Entry, level int) int {
	if len(entries) == 0 {
		return level
	}
	return level
}

// FindParentGroup finds the parent group of the currently selected item
// Returns the index of the parent group in the flattened list, or -1 if not found
func FindParentGroup(entries []Entry, currentIndex int) int {
	flattened := FlattenEntries(entries, 0)
	if currentIndex < 0 || currentIndex >= len(flattened) {
		return -1
	}

	currentEntry := flattened[currentIndex]
	if currentEntry.Level == 0 {
		return -1 // Already at root level
	}

	// Search backwards for a group at the parent level
	targetLevel := currentEntry.Level - 1
	for i := currentIndex - 1; i >= 0; i-- {
		if flattened[i].Type == EntryTypeGroup && flattened[i].Level == targetLevel {
			return i
		}
	}

	return -1
}

// CountVisibleEntries returns the total number of visible entries (respecting collapsed groups)
func CountVisibleEntries(entries []Entry) int {
	return len(FlattenEntries(entries, 0))
}
