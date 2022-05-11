package lesson

import (
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

var temp *template.Template

func ParseLesson(content []byte) Lesson {
	Lesson := new(Lesson)
	parsed := goldmark.DefaultParser().Parse(text.NewReader(content))

	ast.Walk(parsed, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		if !entering || n.Type() == ast.TypeDocument {
			return ast.WalkContinue, nil
		}
		var err error

		if n.Kind() == ast.KindHeading {
			currentHeading := n.(*ast.Heading)
			if currentHeading.Level == 1 {
				Lesson.Name = string(currentHeading.Text(content))
			}
		}

		if n.Kind() == ast.KindParagraph {
			if n.PreviousSibling().Kind() == ast.KindHeading {
				parentHeading := n.PreviousSibling().(*ast.Heading)
				if parentHeading.Level == 1 {
					Lesson.Description = string(n.Text(content))
				}
				if parentHeading.Level == 2 {
					currentTask := new(Task)
					currentTask.Description = string(n.Text(content))
					Lesson.Tasks = append(Lesson.Tasks, *currentTask)
				}
			}
		}
		return s, err
	})
    //TODO: Extract this into Markdown document
	Lesson.Vocabulary = []string{"pwd", "ls", "cd", "whoami", "uname", "echo", "curl", "man", "clear", "less", "vim"}
	return *Lesson
}
