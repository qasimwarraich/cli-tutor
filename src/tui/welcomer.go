package tui

import (
	"fmt"
	"time"

	"cli-tutor/src/lesson"
	"cli-tutor/src/printer"

	"github.com/muesli/termenv"
)

func PrintWelcome(currentLesson lesson.Lesson) {
	termenv.ClearScreen()
	printer.Print("Welcome to Chistole", "welcome")
	time.Sleep(1 * time.Second)
	termenv.ClearScreen()
	printer.Print("This lesson is titled:", "tip")
	printer.Print(currentLesson.Name, "")
	printer.Print("\n"+currentLesson.Description, "guide")
	// TODO: Fix any key bug
	printer.Print("\n\nWhen you are ready press the enter key to begin", "note")
	fmt.Scanln() // Any key
	termenv.ClearScreen()
	printer.Print("Welcome to the shell", "tip")
	printer.Print("Try out some commands or type 'exit'/'quit' to quit the shell", "note")
	time.Sleep(1 * time.Second)
}
