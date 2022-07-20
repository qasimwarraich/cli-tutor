package menuui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 50
	defaultWidth = 80
)

var (
	titleStyle      = lipgloss.NewStyle().MarginLeft(0).Foreground(lipgloss.Color("#34e2e2"))
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(0).PaddingBottom(1)
)

func styleMenu(l *list.Model) {
	// l.SetShowStatusBar(true)
	l.Title = "Which Lesson would you like to start?"
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
}

func styleListItems(d *list.DefaultDelegate) {
	d.Styles.NormalDesc.UnsetWidth()
	d.ShowDescription = true
	d.SetHeight(5)
}

func newList() list.Model {
	items := buildList()
	d := list.NewDefaultDelegate()
	styleListItems(&d)
	l := list.New(items, d, defaultWidth, listHeight)
	styleMenu(&l)
	return l
}
