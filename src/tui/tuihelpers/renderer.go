package tuihelpers

import "github.com/charmbracelet/glamour"

func GetRenderer() *glamour.TermRenderer {
	r, _ := glamour.NewTermRenderer(
		glamour.WithPreservedNewLines(),
        // glamour.WithStylePath("./custom_style.json") TODO: This needs to be embedded into the binary.
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(130),
	)
	return r
}
