# 🧋 bubblebook

> A Storybook-inspired TUI for building and visually testing Bubble Tea components

**bubblebook** is a development tool for Go developers using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework. It provides an interactive terminal-based UI where individual components can be rendered and tested in isolation. This allows for rapid iteration, visual debugging, and easy sharing of component behavior without needing to integrate them into a full application first.

---

## ✨ Features

* 📦 **Component isolation** – Run and inspect individual Bubble Tea components
* 🎛️ **Dynamic knobs** – Modify component props/state on the fly and see results instantly
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

* [ ] Create interface to .Add Datastructures which will Contain the actual component and the knobs and how to put which knob to which param
* [ ] Register & render simple Bubble Tea models
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
