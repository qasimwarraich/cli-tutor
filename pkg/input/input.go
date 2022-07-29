package input

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"cli-tutor/pkg/lesson"
	"cli-tutor/pkg/printer"
	"cli-tutor/pkg/tui/feedback"
)

var previousCommand []string

// Checks to see if given input string is in the vocabulary of the current lesson.
func InputFilter(s string, vocabulary []string) []string {
	input := strings.ReplaceAll(s, "|", " | ")
	split := strings.Fields(input)
	if len(split) > 0 && contains(split[0], vocabulary) {
		// Ensure commands following a pipe are also in the vocabulary.
		if contains("|", split) {
			for i, v := range split {
				if v == "|" {
					if i+1 < len(split) && contains(split[i+1], vocabulary) {
						continue
					} else {
						return []string{}
					}
				}
			}
		}
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
			if fakeJailWarden(filtered_input[1]) {
				err := os.Chdir(filtered_input[1])
				if err != nil {
					printer.Print(err.Error(), "note")
				}
			} else {
				printer.Print(feedback.DangerZone, "error")
			}
		} else {
			err := os.Chdir(os.Getenv("HOME"))
			if err != nil {
				log.Print(err.Error())
			}
		}
	}

	// The double bang needs special treatment as the readline library seems to
	// ignore it.
	if filtered_input[0] == "!!" {
		return RunCommand(previousCommand)
	}
	// Handle piped commands
	if contains("|", filtered_input) {
		command := filtered_input[:]
		commandstring := strings.Join(command, " ")
		out, err := exec.Command("bash", "-c", commandstring).CombinedOutput()
		if err != nil {
			printer.Print(feedback.CommandFailed, "")
			printer.Print(err.Error(), "")
		}
		return string(out)
	}

	if contains("&&", filtered_input) || contains("||", filtered_input) {
		command := filtered_input[:]
		commandstring := strings.Join(command, " ")
		out, err := exec.Command("bash", "-c", commandstring).CombinedOutput()
		if err != nil {
			printer.Print(feedback.CommandFailed, "")
			printer.Print(err.Error(), "")
		}
		return string(out)
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

			printer.Print(feedback.Correct, "guide")
			time.Sleep(2 * time.Second)
			*currentTask++
		} else {
			printer.Print(feedback.Incorrect, "error")
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

	less := exec.Command("less")
	less.Stdin = strings.NewReader(inputstring)
	less.Stdout = os.Stdout

	err := less.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func fakeJailWarden(s string) bool {
	cwd, _ := os.Getwd()
	restricted := []string{
		"/bin", "/boot", "/dev", "/etc", "/home", "/lib", "/lib32",
		"/lib64", "/libx32", "/media", "/mnt", "/opt", "/proc", "/root", "/run", "/sbin",
		"/srv", "/sys", "/tmp", "/usr", "/var",
	}

	if cwd == os.Getenv("HOME") && s == ".." {
		log.Print("Tried to enter danger zone")
		return false
	}

	if s == "/" || s == "/root" {
		return false
	}

	if contains(s, restricted) {
		return false
	}

	return true
}
