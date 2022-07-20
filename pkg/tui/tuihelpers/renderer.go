package tuihelpers

import (
	_ "embed"

	"github.com/charmbracelet/glamour"
)

//go:embed custom_style.json
var stylesheet []byte

func GetRenderer() *glamour.TermRenderer {
	r, _ := glamour.NewTermRenderer(
		glamour.WithPreservedNewLines(),
		glamour.WithStylesFromJSONBytes(stylesheet),
		glamour.WithWordWrap(130),
	)
	return r
}
