package main

import (
	"fmt"
	"os/exec"
	"time"

	"cli-tutor/src/input"
	"cli-tutor/src/lesson"
	"cli-tutor/src/printer"
	"cli-tutor/src/prompt"
	"cli-tutor/src/task"

	"github.com/chzyer/readline"
	"github.com/muesli/termenv"
)

func main() {
	lesson1 := new(lesson.Lesson)

	lesson1.Name = "What is textual interaction"
	lesson1.Vocabulary = []string{"pwd", "ls", "cd", "whoami", "uname", "echo", "curl"}
	lesson1.Description = "This lesson will teach you about the basics of command line interaction"
	lesson1.Tasks = []task.Task{
		{
			Description: "Some spam that will eventually be a task description.",
		},
	}

	termenv.ClearScreen()
	printer.Print("Welcome to Chistole", "welcome")
	time.Sleep(1 * time.Second)

	printer.Print("This lesson is titled:", "")
	printer.Print(lesson1.Name, "")
	printer.Print("\n\n"+lesson1.Description, "")

	printer.Print("\n\nWhen you are ready press the enter key to begin", "note")
	fmt.Scanln() // Any key

	termenv.ClearScreen()

	printer.Print("Welcome to the shell", "")
	printer.Print("Try out some commands or type 'exit'/'quit' to quit the shell", "note")

	rl, err := readline.New(prompt.BuildPrompt() + " > ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	// Readline loop
	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		if line == "" {
			continue
		}

		if line == "quit" || line == "exit" {
			break
		}

		command := input.InputFilter(line, lesson1.Vocabulary)

		var cmd *exec.Cmd

		if len(command) > 0 {
			if len(command) > 1 {
				args := command[1:]
				cmd = exec.Command(command[0], args...)
			} else {
				cmd = exec.Command(line)
			}
			output, _ := cmd.Output()
			fmt.Print(string(output))
		} else {
			printer.Print("Let's stick to the basics", "error")
		}

	}
}
