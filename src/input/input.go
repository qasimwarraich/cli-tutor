package input

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"cli-tutor/src/lesson"
	"cli-tutor/src/printer"
)

var previousCommand []string

// Checks to see if given input string is in the vocabulary of the current lesson.
func InputFilter(s string, vocabulary []string) []string {
	split := strings.Fields(s)
	if len(split) > 0 && contains(split[0], vocabulary) {
		return split
	}
	return []string{}
}

// Takes the now filtered string and runs it as a system call.
func RunCommand(filtered_input []string) string {
	var cmd *exec.Cmd

	//  This is required because cd as a subcommand will not effect the cwd of
	//  the go program.
	if filtered_input[0] == "cd" {
		if len(filtered_input) > 1 {
			err := os.Chdir(filtered_input[1])
			if err != nil {
				printer.Print(err.Error(), "note")
			}
		} else {
			os.Chdir(os.Getenv("$HOME"))
		}
	}

	// The double bang needs special treatment as the readline library seems to
	// ignore it.
	if filtered_input[0] == "!!" {
		return RunCommand(previousCommand)
	}

	// Checks if the command has arguments
	if len(filtered_input) > 1 {
		args := filtered_input[1:]
		cmd = exec.Command(filtered_input[0], args...)
	} else {
		cmd = exec.Command(filtered_input[0])
	}

	// This is needed to handle the !!
	previousCommand = filtered_input

	if filtered_input[0] == "man" {
		out, _ := cmd.CombinedOutput()
		runPager(out)
		return ""
	}
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// Checks if the current lesson task has an expected result and compares the
// result of the user's command to the expected one.
func ValidateCommand(commandOutput string, currentLesson lesson.Lesson, currentTask *int) {
	if currentLesson.Tasks[*currentTask].Expected != "" {
		expected := currentLesson.Tasks[*currentTask].Expected

		if expected[0] == byte('!') {
			expectedcommand := InputFilter(expected[1:], currentLesson.Vocabulary)
			expected = RunCommand(expectedcommand)
		}
		if strings.TrimSpace(commandOutput) == strings.TrimSpace(expected) {

			printer.Print("Yay you did it!, Let's move to the next task.", "guide")
			time.Sleep(2 * time.Second)
			*currentTask++
		} else {
			printer.Print("That isn't quite right", "error")
		}
	}
}

func contains(s string, arr []string) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}
	return false
}

func runPager(input []byte) {
	inputstring := string(input)

	less := exec.Command("less") // TODO: Fix this is docker container.
	less.Stdin = strings.NewReader(inputstring)
	less.Stdout = os.Stdout

	err := less.Run()
	if err != nil {
		log.Fatal(err)
	}
}
