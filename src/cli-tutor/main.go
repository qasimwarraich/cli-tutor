package main

import (
	"log"
	"os"

	"cli-tutor/src/tui"
)

func main() {
	// Init Logging
	logFile, err := os.OpenFile("tutor-log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	tui.StartUI()
	defer os.Remove("file.txt")
}
