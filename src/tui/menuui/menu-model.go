package menuui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 50
	defaultWidth = 20
)

var (
	titleStyle      = lipgloss.NewStyle().MarginLeft(0).Foreground(lipgloss.Color("#34e2e2"))
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(0).PaddingBottom(1)
)

type item struct {
	title       string
	description string
	filename    string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type MenuModel struct {
	list       list.Model
	choice     string
	quitting   bool
	windowsize tea.WindowSizeMsg
}

func (m MenuModel) Init() tea.Cmd {
	// termenv.ClearScreen()
	// printer.Print("Welcome to Chistole", "welcome")
	// time.Sleep(1 * time.Second)
	// termenv.ClearScreen()
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowsize = msg
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 1)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i.filename)
			}
			// return m, tea.Quit
		}
	}
	// switch m.state {
	// case lessonView:
	// 	lmodel := lessonui.New(m.choice)
	// 	// lmodel.Init()
	// 	lmodel.Update(m)
	// case listView:
	// 	m.list.Update(m)
	// }

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MenuModel) View() string {
	return "\n" + m.list.View()
}

func styleMenu(l *list.Model) {
	l.Title = "Which Lesson would you like to start?"
	// l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
}

func New() MenuModel {
	items := []list.Item{
		item{title: "Lesson 1", description: "This lesson is about blah blah", filename: "lesson1.md"},
		item{title: "Lesson 2", description: "This lesson is about spam spam", filename: "lesson2.md"},
	}
	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	styleMenu(&l)

	m := MenuModel{
		list:       l,
		choice:     "",
		quitting:   false,
		windowsize: tea.WindowSizeMsg{},
	}
	return m
}
