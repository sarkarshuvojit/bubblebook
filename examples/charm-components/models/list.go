package models

import (
	"fmt"
	"io"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
)

type item string

func (i item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	if index == m.Index() {
		fmt.Fprint(w, "> "+str)
	} else {
		fmt.Fprint(w, "  "+str)
	}
}

type ListModel struct {
	list list.Model
}

func NewListModel() ListModel {
	items := []list.Item{
		item("Apple"),
		item("Banana"),
		item("Cherry"),
		item("Date"),
		item("Elderberry"),
		item("Fig"),
		item("Grape"),
	}

	l := list.New(items, itemDelegate{}, 30, 10)
	l.Title = "Fruits"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return ListModel{list: l}
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListModel) View() string {
	return "\n" + m.list.View() + "\n\n(Use arrow keys to navigate)\n"
}
