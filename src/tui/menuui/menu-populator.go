package menuui

import (
	"cli-tutor/src/lesson"

	"github.com/charmbracelet/bubbles/list"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func buildList() []list.Item {
	var items []list.Item

	fs := lesson.EmbeddedFS
	dir, _ := fs.ReadDir("lessons")

	for _, file := range dir {
		if file.Type().IsRegular() {
			content, _ := fs.ReadFile("lessons/" + file.Name())
			items = append(items, ParseLesson(file.Name(), content))
		}
	}
	return items
}

func ParseLesson(name string, content []byte) item {
	var item item
	item.filename = name
	parsed := goldmark.DefaultParser().Parse(text.NewReader(content))

	ast.Walk(parsed, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		n.Dump(content, 5)
		s := ast.WalkStatus(ast.WalkContinue)
		if !entering || n.Type() == ast.TypeDocument {
			return ast.WalkContinue, nil
		}
		var err error

		if n.Kind() == ast.KindHeading {
			currentHeading := n.(*ast.Heading)
			if currentHeading.Level == 1 {
				item.title = string(currentHeading.Text(content))
			}
		}

		if n.Kind() == ast.KindParagraph {
			if n.PreviousSibling() != nil && n.PreviousSibling().Kind() == ast.KindHeading {
				parentHeading := n.PreviousSibling().(*ast.Heading)
				if parentHeading.Level == 1 {
					desc := lesson.AssembleLines(n, content)
					item.description = desc
				}

			}
		}
		return s, err
	})
	return item
}
