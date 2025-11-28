# bubblebook

A Storybook-inspired development tool for building and testing Bubble Tea components in isolation.

## Overview

bubblebook provides an interactive terminal UI for developing, testing, and showcasing individual Bubble Tea components without integrating them into a full application. Register your components, preview them in real-time, and iterate quickly with built-in keyboard navigation and focus management.

## Installation

```bash
go get github.com/sarkarshuvojit/bubblebook/pkg/bubblebook
```

## Quick Start

```go
package main

import (
    "github.com/sarkarshuvojit/bubblebook/pkg/bubblebook"
    tea "github.com/charmbracelet/bubbletea"
)

func main() {
    // Register your components
    bubblebook.Register("Button", func() tea.Model {
        return NewButtonModel()
    })

    bubblebook.Register("Input", func() tea.Model {
        return NewInputModel()
    })

    // Launch the interactive TUI
    bubblebook.Start()
}
```

## Usage

### Registering Components

Use `bubblebook.Register()` to add components to the registry. Each component needs a name and a factory function that returns a `tea.Model`:

```go
bubblebook.Register("ComponentName", func() tea.Model {
    return yourComponent{}
})
```

The factory function allows you to:
- Initialize components with specific props
- Set up initial state
- Configure component variants

Example with configuration:

```go
bubblebook.Register("Primary Button", func() tea.Model {
    return NewButton("Click me", PrimaryStyle)
})

bubblebook.Register("Secondary Button", func() tea.Model {
    return NewButton("Cancel", SecondaryStyle)
})
```

### Starting the TUI

After registering your components, launch the bubblebook interface:

```go
bubblebook.Start()
```

This opens a terminal UI where you can:
- Browse all registered components
- Preview components in isolation
- Interact with components in real-time
- Test component behavior and rendering

### Keyboard Controls

- `↑/k`, `↓/j` - Navigate component list
- `g`, `G` - Jump to top/bottom
- `tab` - Switch between list and preview
- `esc` - Return to component list
- `?` - Toggle help screen
- `q`, `ctrl+c` - Quit

## API Reference

### Functions

#### `Register(name string, factory ComponentFactory)`

Adds a component to the bubblebook registry.

**Parameters:**
- `name` - Display name for the component
- `factory` - Function that returns a new instance of `tea.Model`

#### `Start()`

Launches the bubblebook TUI with all registered components.

### Types

#### `ComponentFactory`

```go
type ComponentFactory func() tea.Model
```

A function that returns a fresh instance of a Bubble Tea model.

## Examples

See the [`examples/`](./examples) directory for complete working examples:

- [`examples/basic`](./examples/basic) - Basic usage with spinner, counter, and toggle components
- [`examples/charm-components`](./examples/charm-components) - Examples using Charm libraries

To run an example:

```bash
cd examples/basic
go run main.go
```

## Features

- Component isolation - Run and inspect individual Bubble Tea components
- Pure Bubbletea - Built entirely with Bubbletea, no framework mixing
- Full styling support - Components render with all colors, styles, and ANSI sequences
- Interactive preview - Navigate between components and interact with them in real-time
- Keyboard navigation - Vim-style navigation and intuitive keyboard shortcuts
- Built-in help - Press `?` to see all keyboard shortcuts
- Zero-config - Plug and play with minimal setup

## Use Cases

- **Component Development** - Build and test components in isolation before integration
- **Visual Testing** - Create a visual testing pipeline for TUI components
- **Documentation** - Showcase component variants and states
- **Debugging** - Isolate and debug component behavior
- **Prototyping** - Quickly prototype and iterate on UI elements

## Comparison to Storybook

Just like [Storybook](https://storybook.js.org/) helps frontend developers build and test React/Vue/Svelte components in isolation, bubblebook provides the same workflow for Bubble Tea TUI applications.

## Roadmap

- [x] Register and render Bubble Tea models
- [x] Pure Bubbletea implementation
- [x] Keyboard navigation and focus management
- [x] Built-in help screen
- [ ] Dynamic props via "knobs" (labels, booleans, enums)
- [ ] Live reload on source file change
- [ ] Theming support
- [ ] Export visual snapshots for documentation/testing

## Contributing

Contributions are welcome. Please open an issue or pull request on GitHub.

## License

MIT © 2025 Shuvojit Sarkar
