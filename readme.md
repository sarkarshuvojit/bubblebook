# 🧋 bubblebook

> A Storybook-inspired TUI for building and visually testing Bubble Tea components

**bubblebook** is a development tool for Go developers using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework. It provides an interactive terminal-based UI where individual components can be rendered and tested in isolation. This allows for rapid iteration, visual debugging, and easy sharing of component behavior without needing to integrate them into a full application first.

---

## ✨ Features

* 📦 **Component isolation** – Run and inspect individual Bubble Tea components
* 🎨 **Pure Bubbletea** – Built entirely with Bubbletea, no framework mixing
* 🌈 **Full styling support** – Components render with all colors, styles, and ANSI sequences
* 🎛️ **Interactive preview** – Navigate between components and interact with them in real-time
* ⌨️ **Keyboard navigation** – Vim-style navigation (j/k) and intuitive keyboard shortcuts
* ❓ **Built-in help** – Press `?` to see all keyboard shortcuts
* 🖼️ **Live preview** – View component output as you develop
* 🚀 **Zero-config** – Plug and play with minimal setup
* 🧪 **Visual testing** – Great for building a visual testing pipeline for TUIs

---

## 📸 Preview

*(Coming soon – once the core is implemented)*

---

## 🔧 Getting Started

```bash
go install github.com/sarkarshuvojit/bubblebook@latest
```

Then, in your project:

```bash
bubblebook
```

This will automatically discover your exported Bubble Tea components (more details on structure below) and launch the interface.

### Keyboard Shortcuts

* **↑/k, ↓/j** – Navigate component list
* **g, G** – Jump to top/bottom
* **tab** – Switch between list and preview
* **esc** – Return to component list
* **?** – Toggle help screen
* **q, ctrl+c** – Quit

---

## 🧱 Project Structure

bubblebook expects you to export your components with a consistent interface. Example:

```go
// components/button.go
package components

func NewButton(label string) tea.Model {
    return button{label: label}
}
```

You can register components like so:

```go
// bubblebook/components.go
package main

import (
    "github.com/yourusername/yourapp/components"
    "github.com/yourusername/bubblebook"
)

func main() {
    bubblebook.Register("Button", func() tea.Model {
        return components.NewButton("Click me")
    })

    bubblebook.Start()
}
```

*(More auto-discovery / reflection support planned)*

---

## 📚 Why "bubblebook"?

Just like [Storybook](https://storybook.js.org/) does for frontend React/Vue/Svelte components, **bubblebook** helps you isolate and interactively develop **Bubble Tea** TUI components.

---

## 🛣️ Roadmap

* [x] Register & render simple Bubble Tea models
* [x] Pure Bubbletea implementation (no framework mixing)
* [x] Keyboard navigation and focus management
* [x] Built-in help screen
* [ ] Create interface to add datastructures which will contain the actual component and the knobs
* [ ] Add dynamic props via "knobs" (e.g., labels, booleans, enums)
* [ ] Live reload on source file change
* [ ] Integration with popular Go build tools
* [ ] Theming support
* [ ] Export visual snapshots for documentation/testing

---

## 🤝 Contributing

PRs and ideas welcome! If you're a Bubble Tea enthusiast or terminal UI nerd, come help build the best TUI dev toolchain.

---

## 📄 License

MIT © 2025 \[Shuvojit Sarkar]
