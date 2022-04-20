package printer

import (
	"fmt"
	"github.com/muesli/termenv"
)

func Print(s string, string_type string) {
	p := termenv.ColorProfile()
	styled := termenv.String(s)
	switch string_type {
	case "welcome":
		fmt.Println(styled.Foreground(p.Color("#FF2FF2")))

	case "error":
		fmt.Println(styled.Foreground(p.Color("#FF0000")))

	case "note":
		fmt.Println(styled.Foreground(p.Color("#FFFF00")))

	default:
		fmt.Println(styled.Foreground(p.Color("#00FF00")))
	}
}
