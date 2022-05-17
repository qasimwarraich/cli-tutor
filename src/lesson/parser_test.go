package lesson

import (
	"log"
	"os"
	"reflect"
	"testing"
	"text/template"
)

func TestParseLesson(t *testing.T) {
	temp := template.Must(template.New("lesson1.md").Funcs(FuncMap).ParseFiles("../../lessons/lesson1.md"))
	f, _ := os.Create("expanded.md")
	defer os.Remove(f.Name())

	err := temp.Execute(f, "")
	if err != nil {
		log.Panic(err)
	}

	spam, _ := os.ReadFile("expanded.md")

	result := Lesson{Name: "Basics of Textual Interaction", Vocabulary: []string{"pwd", "ls", "cd", "whoami", "uname", "echo", "curl", "man", "clear", "less", "vim"}, Description: "This lesson will introduce you to the very basics of command line interaction.\nWe will introduce you to the tutor program and how to formulate a command.", Tasks: []Task{{Title: "Introduction", Description: "Welcome to the command line, this may look intimidating but, this  \ntutorial will introduce you to the concepts you need to be familiar  \nwith to thrive at the command line.\n\nWhen you are ready type `n` or `next` and hit enter to continue.", Expected: ""}, {Title: "What is the Shell?", Description: "Textual interaction is a form of interacting with your computer's operating\nsystem using textual input. You, the user, can issue `commands` and the `shell`\nwill interpret them and produce an output.\n\nThe `shell` is a program that wraps your operating system and acts an\nintermediary between the user and the operating system. It manages the user's\ninteraction by accepting input in the form of `commands` and relays output in the\nform of produced output, errors or special shell features like shortcuts.\n```txt\n                       shell───────────────────────┐\n                       │     interpret             │\n┌──────┐               │     ▲     │               │\n│ user ├─command──────►├─────┘     │       ┌────┐  │\n└───▲──┘               │           └──────►│ os │  │\n    │                  │                   └──┬─┘  │\n    └─────output───────┤                      │    │\n                       │◄──output─────────────┘    │\n                       │                           │\n                       └───────────────────────────┘\n\n```\n\nThis graph is an illustration of the interaction cycle between the user and shell.\n\nType `n` or `next` and hit enter to continue.", Expected: ""}, {Title: "Command Formulation", Description: "Commands can generally be supplied to the shell by using the following format:\n\n`command` `--flags` `arguments`\n\ne.g. `wc --lines file.txt `\n\nA `command` is actually a program the shell can run. In the above example the program being called is `wc`, a word counting program.\n\n`--flags` are additional cues that can be provided to the program to alter its behaviour. These can be optional, but almost all command line programs have flags to help produce the output the user wants. In the above example we supply the `--lines` flag to let `wc` know we are only interested in the line count and not the rest of the information the program can produce.\n\nFinally, we have the `argument` which is the input to the program, in the case of our example, a file named `file.txt`.\n\nType `n` or `next` and hit enter to continue.", Expected: ""}, {Title: "Try a Command", Description: "Try it out for yourself!\nType the command `whoami`, hit Enter and look at the result.", Expected: "spam"}, {Title: "Prompt", Description: "This funny looking line below you is called the `prompt`.\n\nIt displays information about the:\n\ncurrent user = `spam`← You!\nhost machine = `spam-linux` ← Your machine\nthe current directory you are in = `/home/spam/Programming/Go/cli-tutor/src/lesson` ← Where on you are on your machine.", Expected: ""}}}

	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want Lesson
	}{
		{
			"Lesson 1 Parse",
			args{spam},
			result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLesson(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLesson() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
