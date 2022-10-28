package logger

import (
	"log"
	"os"
)

var hostName string

func InitLogging() {
	// Init Logging
	logFile, err := os.OpenFile("/tmp/tutor-log.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.Println("Starting new tutor session")
}
