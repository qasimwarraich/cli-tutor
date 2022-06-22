package lessonui

import (
	"io/ioutil"

	"github.com/chzyer/readline"
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

var completer = readline.NewPrefixCompleter(
	readline.PcItem("ls",
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("wc",
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("cd", 
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("cat", 
		readline.PcItemDynamic(listFiles("./")),
	),
	readline.PcItem("next"),
	readline.PcItem("prev"),
	readline.PcItem("exit"),
	readline.PcItem("quit"),
)
