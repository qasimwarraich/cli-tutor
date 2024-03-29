package prompt

import (
	"os"
	"os/user"
	"testing"
)

func TestStylePrompt(t *testing.T) {
	type args struct {
		s            string
		stringStyle string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test styling",
			args{"Test", "magenta"},
			"\033[35mTest\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StylePrompt(tt.args.s, tt.args.stringStyle); got != tt.want {
				t.Errorf("StylePrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildPrompt(t *testing.T) {
	user, _ := user.Current()
	username := StylePrompt(user.Username, "blue")

	host, _ := os.Hostname()
	hostname := StylePrompt(host, "magenta")

	cwd, _ := os.Getwd()
	styledCWD := StylePrompt(cwd, "green")

	tests := []struct {
		name string
		want string
	}{
		{
			"Test Prompt Building",
			username + " @ " + hostname + " in " + styledCWD,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildPrompt(); got != tt.want {
				t.Errorf("BuildPrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}
