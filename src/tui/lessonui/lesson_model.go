package lessonui

import (
	"fmt"
	"log"

	"cli-tutor/src/input"
	"cli-tutor/src/lesson"
	"cli-tutor/src/printer"
	"cli-tutor/src/prompt"
	"cli-tutor/src/tui/tuihelpers"

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

func New(choice string) LessonModel {
	renderer := tuihelpers.GetRenderer()

	// Init Readline
	rl, err := readline.New(prompt.BuildPrompt() + " > ")
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
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, func() tea.Msg {
		w, h, _ := term.GetSize(0)
		return tea.Msg(tea.WindowSizeMsg{w, h}) // HACK: forces a redraw when switching back to menu screen.
	}
}

func (m LessonModel) View() string {
	s := ""

	return s
}

func (m *LessonModel) rline() {
	// Readline loop
	currentTask := 0
	for {
		log.Print(m.rl.Config.Prompt)
		if currentTask < 0 {
			currentTask = 0
		}

		if currentTask >= len(m.currentLesson.Tasks) {
			printer.Print("All Task Completed!, Exiting tutor!", "note")
			break
		}

		tracker := fmt.Sprintf("\n\n%s : %s [%d/%d]:", m.currentLesson.Name, m.currentLesson.Tasks[currentTask].Title, currentTask, len(m.currentLesson.Tasks)-1)
		printer.Print(tracker, "guide")

		out, _ := m.r.Render(m.currentLesson.Tasks[currentTask].Description)
		printer.Print(out, "")

		line, err := m.rl.Readline()
		log.Print(line)
		if err != nil { // io.EOF
			break
		}

		if line == "" {
			continue
		}

		if line == "next" || line == "n" {
			currentTask++
			continue
		}

		if line == "prev" || line == "p" {
			currentTask--
			continue
		}

		if line == "quit" || line == "exit" {
			printer.Print("Exiting tutor, Good bye!", "note")
			m.quitting = true
			break
		}

		filtered_input := input.InputFilter(line, m.currentLesson.Vocabulary)
		if len(filtered_input) == 0 {
			printer.Print("Let's stick to the basics", "error")
			continue
		}

		/* As a final step if the command is in our allowed vocabulary we run
		 * it as a system call, display it's output and validate it against the
		 * expected value on the lesson if it exists. */
		output := input.RunCommand(filtered_input)
		printer.Print(string(output), "")
		input.ValidateCommand(output, m.currentLesson, &currentTask)

	}
}
