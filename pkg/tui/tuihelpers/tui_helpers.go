package tuihelpers

import (
	"os"

	"github.com/muesli/termenv"
)

var output = termenv.NewOutput(os.Stdout)

func ClearScreen() {
	output.ClearScreen()

}
