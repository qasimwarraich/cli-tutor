package tui

import (
	"fmt"
	"os"

	"cli-tutor/src/tui/lessonui"
	"cli-tutor/src/tui/menuui"
	"cli-tutor/src/tui/tuihelpers"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

type sessionState int

const (
	menuView sessionState = iota
	lessonView
)

type MainModel struct {
	state      sessionState
	menu       tea.Model
	lesson     tea.Model
	quitting   bool
	windowsize tea.WindowSizeMsg
}

func (m MainModel) Init() tea.Cmd {
	tuihelpers.ProgramWelcome()
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowsize = msg // pass this along to the entry view so it uses the full window size when it's initialized

	case lessonui.BackMsg:
		m.state = menuView

	case menuui.SelectMessage:
		m.state = lessonView
		m.lesson = lessonui.New(msg.SelectedLesson)

		termenv.ClearScreen()
	}

	switch m.state {
	case menuView:
		menu, cmd := m.menu.Update(msg)
		cmds = append(cmds, cmd)
		menuModel, ok := menu.(menuui.MenuModel)
		if !ok {
			panic("something went wrong with the menu ui ")
		}
		m.menu = menuModel

	case lessonView:
		lesson, cmd := m.lesson.Update(msg)
		cmds = append(cmds, cmd)
		lessonModel, ok := lesson.(lessonui.LessonModel)
		if !ok {
			panic("something went wrong with the lesson ui ")
		}
		m.lesson = lessonModel
	}
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {

	case lessonView:
		return m.lesson.View()

	default:
		return m.menu.View()

	}
}

func New() MainModel {
	return MainModel{
		state:      menuView,
		menu:       menuui.New(),
		lesson:     nil,
		quitting:   false,
		windowsize: tea.WindowSizeMsg{},
	}
}

func StartUI() {
	m := New()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
