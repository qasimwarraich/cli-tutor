package lesson

import (
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func assembleLines(n ast.Node, content []byte) string {
	s := ""
	for i := 0; i < n.Lines().Len(); i++ {
		line := n.Lines().At(i)
		s = s + string(line.Value(content))
	}
	return s
}

func ParseLesson(content []byte) Lesson {
	Lesson := new(Lesson)
	parsed := goldmark.DefaultParser().Parse(text.NewReader(content))

	ast.Walk(parsed, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		if !entering || n.Type() == ast.TypeDocument {
			return ast.WalkContinue, nil
		}
		var err error

        n.Dump(content, 2)
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
					lessonString := assembleLines(n, content)
					Lesson.Description = lessonString
				}
				if parentHeading.Level == 2 {
					currentTask := new(Task)
					taskString := assembleLines(n, content)

					// Handles nested paragraphs
					if n.NextSibling() != nil && n.NextSibling().Kind() != ast.KindHeading {
						for p := n; p.NextSibling().Kind() != ast.KindHeading; p = p.NextSibling() {
							if p.NextSibling().Kind() == ast.KindFencedCodeBlock {
								taskString = taskString +  "\n```txt\n" + assembleLines(p.NextSibling(), content) + "\n```"
							} else {
								taskString = taskString + "\n\n" + assembleLines(p.NextSibling(), content)
							}
						}
					}
					currentTask.Description = taskString
					Lesson.Tasks = append(Lesson.Tasks, *currentTask)
				}
				if parentHeading.Level == 3 {
					vocabularyString := n.Text(content)
					vocabulary := strings.Split(string(vocabularyString), ", ")
					Lesson.Vocabulary = vocabulary
				}
			}
		}
		return s, err
	})
	return *Lesson
}
