package tuihelpers

import "github.com/charmbracelet/glamour"

func GetRenderer() *glamour.TermRenderer {
	r, _ := glamour.NewTermRenderer(
		glamour.WithPreservedNewLines(),
		glamour.WithStylePath("./custom_style.json"),
		glamour.WithWordWrap(130),
	)
	return r
}
