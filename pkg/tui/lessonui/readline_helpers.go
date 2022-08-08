package lessonui

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"

	"cli-tutor/pkg/printer"

	"github.com/chzyer/readline"
	"github.com/muesli/termenv"
)

func listFiles(path string) func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

func listDirs(path string) func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			if f.IsDir() {
				names = append(names, f.Name())
			}
		}
		return names
	}
}

// Lists words in the vocabulary of the current lesson.
func listVocabulary() func(string) []string {
	return func(line string) []string {
		return CurrentLesson.Vocabulary
	}
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("ls",
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("wc",
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("cd",
		readline.PcItemDynamic(listDirs("./")),
	),
	readline.PcItem("cat",
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("rm",
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("rmdir",
		readline.PcItemDynamic(listDirs("./")),
	),
	// TODO: This breaks the rest of the completions above
	// readline.PcItemDynamic(listVocabulary()),

	readline.PcItem("next"),
	readline.PcItem("prev"),
	readline.PcItem("exit"),
	readline.PcItem("quit"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func ToggleZen() {
	if ZenMode == false {
		termenv.ClearScreen()
		printer.Print("zen mode toggeled on", "note")
	} else {
		printer.Print("zen mode toggeled off", "note")
	}

	ZenMode = !ZenMode
}

func ZenPrint(input, output, prompt string, style string) {
	termenv.ClearScreen()
	fmt.Println(prompt+input, "")
	printer.Print(output, style)
}

func GetLessonNumber(s string) (int, error) {
	r := regexp.MustCompile("[0-9]+")
	nums := r.FindAllString(s, -1)

	return strconv.Atoi(nums[0])
}
