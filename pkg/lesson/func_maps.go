package lesson

import (
	"os"
	"os/user"
	"text/template"

	"cli-tutor/pkg/logger"
)

func testFunc() string {
	return "TEST"
}

func getUser() string {
	if logger.Identifier != "" {
		return logger.Identifier
	} else {
		current_user, _ := user.Current()
		return current_user.Username
	}
}

func getHost() string {
	hostname, _ := os.Hostname()
	return hostname
}

func getWd() string {
	cwd, _ := os.Getwd()
	return cwd
}

var FuncMap = template.FuncMap{
	"TestFunc": testFunc,
	"GetUser":  getUser,
	"GetHost":  getHost,
	"GetWd":    getWd,
}
