package styler

import (

	"github.com/muesli/termenv"
)

func Style(s string, string_type string) termenv.Style {
	p := termenv.ColorProfile()
	styled := termenv.String(s)
	switch string_type {
	case "welcome":
		return styled.Foreground(p.Color("#FF2FF2"))
        

	case "error":
		return styled.Foreground(p.Color("#FF0000"))

	case "note":
		return styled.Foreground(p.Color("#FFFF00"))

	case "tip":
        return  styled.Foreground(p.Color("#00FFF0"))
    
    case "guide":
		return styled.Foreground(p.Color("#00FF00"))

	default:
		return styled
	}
}
