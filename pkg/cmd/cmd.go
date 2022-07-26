package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"cli-tutor/pkg/logger"
	"cli-tutor/pkg/printer"
	"cli-tutor/pkg/tui"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-tutor",
	Short: "cli-tutor is an application that introduces you to the command line.",
	Long: `
        _ _       __        __
  _____/ (_)     / /___  __/ /_____  _____
 / ___/ / /_____/ __/ / / / __/ __ \/ ___/
/ /__/ / /_____/ /_/ /_/ / /_/ /_/ / /    
\___/_/_/      \__/\__,_/\__/\____/_/     
                                        
A simple command line tutor application that aims to introduce users to the 
    basics of command line interaction.
    Web version is available at https://tutor.chistole.ch`,

	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogging()

		nouploadflag, err := cmd.Flags().GetBool("no-upload-log")
		if err != nil {
			log.Panicln(err)
		}
		if !nouploadflag {

			printer.Print(`
You have started the tutor with the option of sending a log of your tutor
session to the developer. To opt-out of this mode supply the name 'exit', hit
enter and restart the tool providing the -n or --no-upload-log file. 

For more info type 'cli-tutor info' after exiting.
            `, "tip")
			for logger.Identifier == "" {
				rl, _ := readline.New("Please enter a name or id to identify your log file or enter 'exit' > ")
				line, _ := rl.Readline()
				line = strings.TrimSpace(line)
				if line == "exit" {
					os.Exit(0)
				}
				logger.Identifier = line
			}
			defer logger.UploadLog()
		}

		tui.StartUI()
		defer os.Remove("file.txt")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli-tutor",
	Long:  `All software has versions. This is cli-tutor's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-tutor v0.0 -- HEAD")
	},
}

func init() {
	rootCmd.Flags().BoolP("no-upload-log", "n", false, "Do not send a copy of the log to the developer")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
