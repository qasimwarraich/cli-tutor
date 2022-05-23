package main

import (
	"fmt"
	"log"
	"os"

	"cli-tutor/src/input"
	"cli-tutor/src/lesson"
	"cli-tutor/src/printer"
	"cli-tutor/src/prompt"
	"cli-tutor/src/tui"

	"github.com/chzyer/readline"
)

func main() {
	// Init Logging
	logFile, err := os.OpenFile("tutor-log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	// Load lesson and create example lesson file
	currentLesson := lesson.LoadLesson()
	defer os.Remove("file.txt")
	tui.PrintWelcome(currentLesson)
	r := tui.GetRenderer()

	// Init Readline
	rl, err := readline.New(prompt.BuildPrompt() + " > ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	// Readline loop
	currentTask := 0
	for {
		log.Print(rl.Config.Prompt)
		if currentTask < 0 {
			currentTask = 0
		}

		if currentTask >= len(currentLesson.Tasks) {
			printer.Print("All Task Completed!, Exiting tutor!", "note")
			break
		}

		tracker := fmt.Sprintf("\n\n%s : %s [%d/%d]:", currentLesson.Name, currentLesson.Tasks[currentTask].Title, currentTask, len(currentLesson.Tasks)-1)
		printer.Print(tracker, "guide")

		out, _ := r.Render(currentLesson.Tasks[currentTask].Description)
		printer.Print(out, "")

		line, err := rl.Readline()
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
			break
		}

		filtered_input := input.InputFilter(line, currentLesson.Vocabulary)
		if len(filtered_input) == 0 {
			printer.Print("Let's stick to the basics", "error")
			continue
		}

		/* As a final step if the command is in our allowed vocabulary we run
		 * it as a system call, display it's output and validate it against the
		 * expected value on the lesson if it exists. */
		output := input.RunCommand(filtered_input)
		printer.Print(string(output), "")
		input.ValidateCommand(output, currentLesson, &currentTask)

	}
}
