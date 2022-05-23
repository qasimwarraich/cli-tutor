package input

import (
	"os/exec"
	"strings"
	"time"

	"cli-tutor/src/lesson"
	"cli-tutor/src/printer"
)

func InputFilter(s string, vocabulary []string) []string {
	split := strings.Fields(s)
	if len(split) > 0 && contains(split[0], vocabulary) {
		return split
	}
	return []string{}
}

func RunCommand(filtered_input []string) string {
	var cmd *exec.Cmd

	if len(filtered_input) > 1 {
		args := filtered_input[1:]
		cmd = exec.Command(filtered_input[0], args...)
	} else {
		cmd = exec.Command(filtered_input[0])
	}
	output, _ := cmd.CombinedOutput()
	return string(output)
}

func ValidateCommand(commandOutput string, currentLesson lesson.Lesson, currentTask *int) {
	if currentLesson.Tasks[*currentTask].Expected != "" {
		if strings.TrimSpace(commandOutput) == currentLesson.Tasks[*currentTask].Expected {
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
