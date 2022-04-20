package main

import (
	"fmt"
	"os/exec"
	"time"

	"cli-tutor/src/printer"
	"cli-tutor/src/prompt"
	"cli-tutor/src/input"

	"github.com/chzyer/readline"
	"github.com/muesli/termenv"
)




func main() {
	rl, err := readline.New(prompt.BuildPrompt() + " > ")

	if err != nil {
		panic(err)
	}
	defer rl.Close()

	termenv.ClearScreen()

	printer.Print("Welcome to Chistole", "welcome")

	time.Sleep(1 * time.Second)

	printer.Print("This lesson will be about", "")
	printer.Print(`Lorem ipsum dolor sit amet, officia excepteur ex fugiat
    reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse
    exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit
    nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor
    minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure
    elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor
    Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing
    id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua
    reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate
    laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa
    duis.`, "")

	// time.Sleep(4 * time.Second)

	printer.Print("\n\nWhen you are ready press any key to begin", "note")
	fmt.Scanln() // Any key

	termenv.ClearScreen()

	printer.Print("Welcome to the shell", "")
	printer.Print(`Try out some commands or type "exit"/"quit" to quit the shell`, "note")

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

        printer.Print("spam", "")
		command := input.InputFilter(line)

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
