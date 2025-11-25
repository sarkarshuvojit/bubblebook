package bubblebook

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// ComponentFactory is a function that returns a new instance of a Bubble Tea model.
type ComponentFactory func() tea.Model

type componentEntry struct {
	name    string
	factory ComponentFactory
}

var components []componentEntry

// Register adds a component to the Bubblebook registry.
func Register(name string, factory ComponentFactory) {
	components = append(components, componentEntry{name: name, factory: factory})
}

// tviewWriter bridges Bubble Tea output to a tview TextView.
type tviewWriter struct {
	app *tview.Application
	tv  *tview.TextView
}

func (w tviewWriter) Write(p []byte) (int, error) {
	s := string(p)
	// Strip ANSI escape sequences that Bubble Tea uses for cursor control
	// This prevents raw escape codes like "[2A" from appearing in the display
	s = stripAnsiEscapes(s)
	w.app.QueueUpdateDraw(func() {
		w.tv.SetText(s)
	})
	return len(p), nil
}

// stripAnsiEscapes removes ANSI escape sequences from a string
func stripAnsiEscapes(s string) string {
	// Simple approach: remove ESC[ sequences
	// ESC is \x1b, followed by [, then any number of parameter bytes, then a command byte
	result := ""
	i := 0
	for i < len(s) {
		if i < len(s)-1 && s[i] == '\x1b' && s[i+1] == '[' {
			// Found escape sequence, skip until we find the command byte
			i += 2
			for i < len(s) && ((s[i] >= '0' && s[i] <= '9') || s[i] == ';' || s[i] == '?') {
				i++
			}
			// Skip the command byte
			if i < len(s) {
				i++
			}
		} else if i < len(s)-1 && s[i] == '\x1b' && s[i+1] == ']' {
			// OSC sequence (e.g., for setting title)
			i += 2
			for i < len(s) && s[i] != '\x07' && s[i] != '\x1b' {
				i++
			}
			if i < len(s) && s[i] == '\x07' {
				i++
			} else if i < len(s)-1 && s[i] == '\x1b' && s[i+1] == '\\' {
				i += 2
			}
		} else {
			result += string(s[i])
			i++
		}
	}
	return result
}


// convertTcellKey converts a tcell EventKey to a Bubble Tea KeyMsg.
func convertTcellKey(ev *tcell.EventKey) tea.KeyMsg {
	key := ev.Key()
	runes := ev.Rune()
	mod := ev.Modifiers()

	// Build the tea.KeyMsg
	var keyMsg tea.KeyMsg

	// Handle special keys
	switch key {
	case tcell.KeyEnter:
		keyMsg.Type = tea.KeyEnter
	case tcell.KeyTab:
		keyMsg.Type = tea.KeyTab
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		keyMsg.Type = tea.KeyBackspace
	case tcell.KeyEscape:
		keyMsg.Type = tea.KeyEscape
	case tcell.KeyUp:
		keyMsg.Type = tea.KeyUp
	case tcell.KeyDown:
		keyMsg.Type = tea.KeyDown
	case tcell.KeyLeft:
		keyMsg.Type = tea.KeyLeft
	case tcell.KeyRight:
		keyMsg.Type = tea.KeyRight
	case tcell.KeyHome:
		keyMsg.Type = tea.KeyHome
	case tcell.KeyEnd:
		keyMsg.Type = tea.KeyEnd
	case tcell.KeyPgUp:
		keyMsg.Type = tea.KeyPgUp
	case tcell.KeyPgDn:
		keyMsg.Type = tea.KeyPgDown
	case tcell.KeyDelete:
		keyMsg.Type = tea.KeyDelete
	case tcell.KeyInsert:
		keyMsg.Type = tea.KeyInsert
	case tcell.KeyRune:
		// Regular character
		keyMsg.Type = tea.KeyRunes
		keyMsg.Runes = []rune{runes}
	default:
		// For other keys, try to use the rune if available
		if runes != 0 {
			keyMsg.Type = tea.KeyRunes
			keyMsg.Runes = []rune{runes}
		}
	}

	// Handle modifiers
	if mod&tcell.ModAlt != 0 {
		keyMsg.Alt = true
	}

	return keyMsg
}

// Start launches the Bubblebook TUI.
func Start() {
	app := tview.NewApplication()

	// UI Elements
	list := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedBackgroundColor(tcell.ColorDarkSlateBlue)
	list.SetTitle(" Components ").SetBorder(true)

	mainArea := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false) // Bubble Tea handles wrapping usually
	mainArea.SetTitle(" Preview ").SetBorder(true)

	// Layout
	flex := tview.NewFlex().
		AddItem(list, 25, 1, true).
		AddItem(mainArea, 0, 4, false)

	// State for the currently running Bubble Tea program
	var currentProgram *tea.Program
	var programKilled bool // Track if we intentionally killed the program

	// Function to load a component
	loadComponent := func(index int) {
		if index < 0 || index >= len(components) {
			return
		}

		// Stop previous program if running
		if currentProgram != nil {
			programKilled = true // Mark that we're intentionally killing it
			currentProgram.Kill()
			currentProgram = nil
		}

		entry := components[index]
		model := entry.factory()

		// Create a writer that updates the mainArea
		w := tviewWriter{app: app, tv: mainArea}

		// Start the new program with input disabled (we'll forward manually)
		programKilled = false // Reset the flag for the new program
		currentProgram = tea.NewProgram(model, tea.WithOutput(w), tea.WithInput(nil))
		
		go func() {
			prog := currentProgram // Capture the current program
			if _, err := prog.Run(); err != nil {
				// Only show error if we didn't intentionally kill the program
				if !programKilled {
					w.Write([]byte(fmt.Sprintf("Error running component: %v", err)))
				}
			}
		}()
	}

	// Set up input capture for mainArea to forward events to the Bubble Tea program
	mainArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			// Return focus to list
			app.SetFocus(list)
			return nil
		}

		if currentProgram != nil {
			// Convert and send the key event to the Bubble Tea program
			keyMsg := convertTcellKey(event)
			currentProgram.Send(keyMsg)
		}

		// Consume the event (don't pass to tview)
		return nil
	})

	// Populate list
	for i, c := range components {
		// Capture index for closure
		idx := i
		list.AddItem(c.name, "", 0, func() {
			loadComponent(idx)
			app.SetFocus(mainArea) // Focus the component when selected
		})
	}

	// Handle selection change (optional: auto-load on navigation?)
	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		loadComponent(index)
	})
	
	// Select first component if available
	if len(components) > 0 {
		loadComponent(0)
	}

	if err := app.SetRoot(flex, true).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running bubblebook: %v\n", err)
		os.Exit(1)
	}
}
