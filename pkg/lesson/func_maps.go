package lesson

import (
	"os"
	"os/user"
	"text/template"
)

func testFunc() string {
	return "TEST"
}

func getUser() string {
	current_user, _ := user.Current()
	username := current_user.Username
	return username
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
