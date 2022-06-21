package tuihelpers

import (
	"os"
	"time"

	"cli-tutor/src/lesson"
	"cli-tutor/src/printer"

	"github.com/muesli/termenv"
)

func PrintWelcome(currentLesson lesson.Lesson) {
	printer.Print("This lesson is titled:", "tip")
	printer.Print(currentLesson.Name, "")
	printer.Print("\n"+currentLesson.Description, "guide")
	printer.Print("\n\nWhen you are ready press any key to begin", "note")
	b := make([]byte, 1) // Any Key
	os.Stdin.Read(b)
	termenv.ClearScreen()
	printer.Print("Welcome to the shell", "tip")
	printer.Print("Try out some commands or type 'exit'/'quit' to quit the shell", "note")
	time.Sleep(1 * time.Second)
}
