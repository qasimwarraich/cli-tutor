package lessonui

import (
	"log"

	"cli-tutor/pkg/lesson"
	"cli-tutor/pkg/tui/tuihelpers"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/chzyer/readline"
	"github.com/muesli/termenv"
	"golang.org/x/term"
)

type BackMsg bool

type LessonModel struct {
	currentLesson lesson.Lesson
	rl            *readline.Instance
	r             *glamour.TermRenderer
	quitting      bool
}

func (m LessonModel) GetCurrentLesson() *lesson.Lesson { return &m.currentLesson }

var CurrentLesson *lesson.Lesson

func New(choice string) LessonModel {
	renderer := tuihelpers.GetRenderer()
	lessonNumber, err := GetLessonNumber(choice)
	if err != nil {
		log.Println(err)
	}

	if lessonNumber > 1 {
		ZenMode = false
	}

	// Init Readline
	rl, err := readline.NewEx(&readline.Config{
		Prompt:              ">",
		HistoryFile:         "./.tutor_history",
		HistoryLimit:        999,
		InterruptPrompt:     "^C",
		AutoComplete:        completer,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	return LessonModel{
		currentLesson: lesson.LoadLesson(choice),
		rl:            rl,
		r:             renderer,
		quitting:      false,
	}
}

func (m LessonModel) Init() tea.Cmd {
	return nil
}

func (m LessonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	CurrentLesson = m.GetCurrentLesson()
	if !m.quitting {
		m.rline()
	} else {
		m.rl.Close()
		termenv.ClearScreen()
		return m, func() tea.Msg {
			return BackMsg(true)
		}
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// case "ctrl+c":
		// return m, tea.Quit
		}
	}

	return m, func() tea.Msg {
		w, h, _ := term.GetSize(0)
		return tea.Msg(tea.WindowSizeMsg{Width: w, Height: h}) // HACK: forces a redraw when switching back to menu screen.
	}
}

func (m LessonModel) View() string {
	return ""
}
