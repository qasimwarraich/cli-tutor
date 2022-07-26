package tuihelpers

import (
	"os"
	"time"

	"cli-tutor/pkg/lesson"
	"cli-tutor/pkg/printer"

	"github.com/muesli/termenv"
)

func LessonWelcome(currentLesson lesson.Lesson) {
	printer.Print("This lesson is titled:", "tip")
	printer.Print(currentLesson.Name, "")
	printer.Print("\n"+currentLesson.Description, "guide")
	printer.Print("\n\nWhen you are ready press any key to begin", "note")
	b := make([]byte, 1) // Any Key
	os.Stdin.Read(b)
	termenv.ClearScreen()
	printer.Print("Welcome to the shell", "tip")
	printer.Print("Type 'commands' to view a list of available commands or type 'exit'/'quit' to quit the shell", "note")
	time.Sleep(1 * time.Second)
}

func ProgramWelcome() {
	s := `
          ___           ___                       ___           ___           ___           ___       ___     
         /\  \         /\__\          ___        /\  \         /\  \         /\  \         /\__\     /\  \    
        /::\  \       /:/  /         /\  \      /::\  \        \:\  \       /::\  \       /:/  /    /::\  \   
       /:/\:\  \     /:/__/          \:\  \    /:/\ \  \        \:\  \     /:/\:\  \     /:/  /    /:/\:\  \  
      /:/  \:\  \   /::\  \ ___      /::\__\  _\:\~\ \  \       /::\  \   /:/  \:\  \   /:/  /    /::\~\:\  \ 
     /:/__/ \:\__\ /:/\:\  /\__\  __/:/\/__/ /\ \:\ \ \__\     /:/\:\__\ /:/__/ \:\__\ /:/__/    /:/\:\ \:\__\
     \:\  \  \/__/ \/__\:\/:/  / /\/:/  /    \:\ \:\ \/__/    /:/  \/__/ \:\  \ /:/  / \:\  \    \:\~\:\ \/__/
      \:\  \            \::/  /  \::/__/      \:\ \:\__\     /:/  /       \:\  /:/  /   \:\  \    \:\ \:\__\  
       \:\  \           /:/  /    \:\__\       \:\/:/  /     \/__/         \:\/:/  /     \:\  \    \:\ \/__/  
        \:\__\         /:/  /      \/__/        \::/  /                     \::/  /       \:\__\    \:\__\    
         \/__/         \/__/                     \/__/                       \/__/         \/__/     \/__/    

      ___           ___                            ___           ___           ___           ___           ___     
     /\  \         /\__\      ___                 /\  \         /\__\         /\  \         /\  \         /\  \    
    /::\  \       /:/  /     /\  \                \:\  \       /:/  /         \:\  \       /::\  \       /::\  \   
   /:/\:\  \     /:/  /      \:\  \                \:\  \     /:/  /           \:\  \     /:/\:\  \     /:/\:\  \  
  /:/  \:\  \   /:/  /       /::\__\               /::\  \   /:/  /  ___       /::\  \   /:/  \:\  \   /::\~\:\  \ 
 /:/__/ \:\__\ /:/__/     __/:/\/__/              /:/\:\__\ /:/__/  /\__\     /:/\:\__\ /:/__/ \:\__\ /:/\:\ \:\__\
 \:\  \  \/__/ \:\  \    /\/:/  /                /:/  \/__/ \:\  \ /:/  /    /:/  \/__/ \:\  \ /:/  / \/_|::\/:/  /
  \:\  \        \:\  \   \::/__/                /:/  /       \:\  /:/  /    /:/  /       \:\  /:/  /     |:|::/  / 
   \:\  \        \:\  \   \:\__\                \/__/         \:\/:/  /     \/__/         \:\/:/  /      |:|\/__/  
    \:\__\        \:\__\   \/__/                               \::/  /                     \::/  /       |:|  |    
     \/__/         \/__/                                        \/__/                       \/__/         \|__|    
     `

	printer.Print(s, "welcome")
	time.Sleep(2 * time.Second)
	termenv.ClearScreen()
}
