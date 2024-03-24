package printer

import (
	"fmt"
    "log"

    "cli-tutor/pkg/colours"

	"github.com/muesli/termenv"
)

func Print(s string, string_type string) {
	p := termenv.ColorProfile()
	styled := termenv.String(s)
	switch string_type {
	case "welcome":
		fmt.Println(styled.Foreground(p.Color(colours.DefaultWelcome)))
		log.Println(styled.Foreground(p.Color(colours.DefaultWelcome)))

	case "error":
		fmt.Println(styled.Foreground(p.Color(colours.DefaultError)))
		log.Println(styled.Foreground(p.Color(colours.DefaultError)))

	case "note":
		fmt.Println(styled.Foreground(p.Color(colours.DefaultNote)))
		log.Println(styled.Foreground(p.Color(colours.DefaultNote)))

	case "tip":
		fmt.Println(styled.Foreground(p.Color(colours.DefaultTip)))
		log.Println(styled.Foreground(p.Color(colours.DefaultTip)))
    
    case "guide":
		fmt.Println(styled.Foreground(p.Color(colours.DefaultGuide)))
		log.Println(styled.Foreground(p.Color(colours.DefaultGuide)))

	default:
		fmt.Println(s)
		log.Println(s)
	}
}
