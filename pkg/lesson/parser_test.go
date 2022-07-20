package lesson

import (
	"log"
	"os"
	"reflect"
	"testing"
	"text/template"
)

func TestParseLesson(t *testing.T) {
	temp := template.Must(template.New("test.md").Funcs(FuncMap).ParseFiles("./test/test.md"))
	f, _ := os.Create("expanded.md")
	defer os.Remove(f.Name())

	err := temp.Execute(f, "")
	if err != nil {
		log.Panic(err)
	}

	expandedTestFile, _ := os.ReadFile("expanded.md")

	result := Lesson{Name: "Lesson Title", Vocabulary: []string{"TEST", "echo"}, Description: "This should get parsed as the lesson description.", Tasks: []Task{{Title: "First Task", Description: "Will this be parsed.\nHopefully!\n\nWill this nested part get parsed?\n\nWill this `code` render.\nand what about this expansion: TEST", Expected: ""}, {Title: "Does the graph work?", Description: "Does it (This title paragraph is necessary)\n```txt\n                       shell───────────────────────┐\n                       │     interpret             │\n┌──────┐               │     ▲     │               │\n│ user ├─command──────►├─────┘     │       ┌────┐  │\n└───▲──┘               │           └──────►│ os │  │\n    │                  │                   └──┬─┘  │\n    └─────output───────┤                      │    │\n                       │◄──output─────────────┘    │\n                       │                           │\n                       └───────────────────────────┘\n\n```\n\nDoes this following paragraph also render?", Expected: ""}, {Title: "Check Expected", Description: "Tests if expected value is parsed correctly. Type `echo TEST`", Expected: "TEST"}}}

	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want Lesson
	}{
		{
			"Parser Test: test/test.md",
			args{expandedTestFile},
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
