package lessonui

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"cli-tutor/src/input"
	"cli-tutor/src/printer"
	"cli-tutor/src/prompt"
	"cli-tutor/src/tui/tuihelpers"

	"github.com/chzyer/readline"
)

func (m *LessonModel) rline() {
	tuihelpers.LessonWelcome(m.currentLesson)

	// os.Chdir("TESTSPACE")
	// unix.Chroot("TESTSPACE")

	// Readline loop
	currentTask := 0
	for {
		m.rl.SetPrompt(prompt.BuildPrompt() + " $ ")
		log.Print(m.rl.Config.Prompt)
		if currentTask < 0 {
			currentTask = 0
		}

		if currentTask >= len(m.currentLesson.Tasks) {
			printer.Print("Congratulations! All Task Completed!, Exiting Lesson!", "note")
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
			break
		}

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
