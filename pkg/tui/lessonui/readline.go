package lessonui

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"cli-tutor/pkg/input"
	"cli-tutor/pkg/logger"
	"cli-tutor/pkg/printer"
	"cli-tutor/pkg/prompt"
	"cli-tutor/pkg/tui/feedback"
	"cli-tutor/pkg/tui/tuihelpers"

	"github.com/chzyer/readline"
	"github.com/muesli/termenv"
)

func (m *LessonModel) rline() {
	tuihelpers.LessonWelcome(m.currentLesson)

	termenv.ShowCursor()

	// Readline loop
	currentTask := 0
	for {
		currentPrompt := prompt.BuildPrompt(logger.Identifier) + " $ "
		m.rl.SetPrompt(currentPrompt)
		log.Print(currentPrompt)
		if currentTask < 0 {
			currentTask = 0
		}

		if currentTask >= len(m.currentLesson.Tasks) {
			printer.Print(feedback.Completed, "note")
			time.Sleep(2 * time.Second)
			m.quitting = true
			break
		}

		tracker := fmt.Sprintf("\n\n%s : %s [%d/%d]:", m.currentLesson.Name, m.currentLesson.Tasks[currentTask].Title, currentTask, len(m.currentLesson.Tasks)-1)
		printer.Print(tracker, "guide")

		out, _ := m.r.Render(m.currentLesson.Tasks[currentTask].Description)
		printer.Print(out, "")

		line, err := m.rl.Readline()
		line = strings.TrimSpace(line)

		if err == readline.ErrInterrupt {
			if len(line) >= 0 {
				continue
			}
		} else if err == io.EOF {
			continue // essentially captures ^D
		}

		log.Print(line)
		if err != nil { // io.EOF
			break
		}

		if line == "" {
			continue
		}

		if line == "next" || line == "n" {
			if m.currentLesson.Tasks[currentTask].Expected != "" {
				printer.Print(feedback.Expected, "error")
			} else {
				currentTask++
			}
			continue
		}

		if line == "prev" || line == "p" {
			currentTask--
			continue
		}

		if line == "commands" {
			vocabulary := fmt.Sprintln(m.currentLesson.Vocabulary)
			printer.Print("Available commands: "+vocabulary, "")
			continue
		}

		if line == "quit" || line == "exit" {
			printer.Print(feedback.Exiting, "note")
			m.quitting = true
			break
		}

		filtered_input := input.InputFilter(line, m.currentLesson.Vocabulary)
		if len(filtered_input) == 0 {
			printer.Print(feedback.InvalidCommand, "error")
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
