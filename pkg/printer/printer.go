package printer

import (
	"fmt"
    "log"

	"github.com/muesli/termenv"
)

func Print(s string, string_type string) {
	p := termenv.ColorProfile()
	styled := termenv.String(s)
	switch string_type {
	case "welcome":
		fmt.Println(styled.Foreground(p.Color("#FF2FF2")))
		log.Println(styled.Foreground(p.Color("#FF2FF2")))

	case "error":
		fmt.Println(styled.Foreground(p.Color("#FF0000")))
		log.Println(styled.Foreground(p.Color("#FF0000")))

	case "note":
		fmt.Println(styled.Foreground(p.Color("#FFFF00")))
		log.Println(styled.Foreground(p.Color("#FFFF00")))

	case "tip":
		fmt.Println(styled.Foreground(p.Color("#00FFF0")))
		log.Println(styled.Foreground(p.Color("#00FFF0")))
    
    case "guide":
		fmt.Println(styled.Foreground(p.Color("#00FF00")))
		log.Println(styled.Foreground(p.Color("#00FF00")))

	default:
		fmt.Println(s)
		log.Println(s)
	}
}
