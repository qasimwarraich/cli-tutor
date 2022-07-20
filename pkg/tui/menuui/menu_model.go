package menuui

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title       string
	description string
	filename    string
}

type SelectMessage struct {
	SelectedLesson string
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
				log.Printf("Selected lesson: %s", m.choice)
				return m, func() tea.Msg {
					return SelectMessage{SelectedLesson: m.choice}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MenuModel) View() string {
	return "\n" + m.list.View()
}

func New() MenuModel {
	m := MenuModel{
		list:       newList(),
		choice:     "",
		quitting:   false,
		windowsize: tea.WindowSizeMsg{},
	}
	return m
}
