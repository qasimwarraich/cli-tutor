package tuihelpers

import "github.com/charmbracelet/glamour"

func GetRenderer() *glamour.TermRenderer {
	r, _ := glamour.NewTermRenderer(
		glamour.WithPreservedNewLines(),
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(130),
	)
	return r
}
