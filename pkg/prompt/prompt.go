package prompt

import (
	"os"
	"os/user"
)

func StylePrompt(s string, stringStyle string) string {
	const (
		colorReset   = "\033[0m"
		colorRed     = "\033[31m"
		colorGreen   = "\033[32m"
		colorYellow  = "\033[33m"
		colorBlue    = "\033[34m"
		colorMagenta = "\033[35m"
		colorCyan    = "\033[36m"
		colorWhite   = "\033[37m"
	)
	switch stringStyle {

	case "magenta":
		return colorMagenta + s + colorReset

	case "red":
		return colorRed + s + colorReset

	case "yellow":
		return colorYellow + s + colorReset

	case "blue":
		return colorBlue + s + colorReset

	default:
		return colorGreen + s + colorReset
	}
}

func BuildPrompt() string {
	var username string

	user, _ := user.Current()
	username = StylePrompt(user.Username, "blue")

	host, _ := os.Hostname()
	hostname := StylePrompt(host, "magenta")

	cwd, _ := os.Getwd()
	styledCwd := StylePrompt(cwd, "green")

	return username + " @ " + hostname + " in " + styledCwd
}
