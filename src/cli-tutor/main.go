package main

import (
	"os"

	"cli-tutor/src/logger"
	"cli-tutor/src/tui"
)

func main() {

    logger.InitLogging()

	tui.StartUI()
	defer os.Remove("file.txt")
	defer logger.UploadLog()
}
