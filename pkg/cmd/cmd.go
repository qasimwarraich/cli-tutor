package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"cli-tutor/pkg/logger"
	"cli-tutor/pkg/tui"
	"cli-tutor/pkg/tui/tuihelpers"

	"github.com/chzyer/readline"
	"github.com/muesli/termenv"
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
			log.Println(err)
		}

		if !nouploadflag {
			for logger.Identifier == "" {
				p := termenv.ColorProfile()
				styled := termenv.String("Please enter a name to identify yourself or enter 'exit' > ")
				prompt := fmt.Sprint(styled.Foreground(p.Color("#01FFF0")))
				rl, _ := readline.New(prompt)
				line, _ := rl.Readline()
				line = strings.TrimSpace(line)
				if line == "exit" || line == "quit" {
					os.Exit(0)
				}
				logger.Identifier = line
                log.Printf("Identifier: %s", logger.Identifier)
			}
			defer logger.UploadLog()
		}

		nowelcomeflag, err := cmd.Flags().GetBool("no-welcome")
		if err != nil {
			log.Println(err)
		}

		if !nowelcomeflag {
			tuihelpers.ProgramWelcome()
		} else {
			termenv.ClearScreen()
		}

		initial_wd, _ := os.Getwd()

		tui.StartUI()

		defer os.Remove(initial_wd + "/file.txt")
		defer os.Remove(initial_wd + "/.hiddenfile.txt")
		defer termenv.ClearScreen()
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

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Prints information about the tool and log collection",
	Long:  `Prints out information regarding the tool in general and it's logging feature`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			`
cli-tutor is a simple command line tutor application that aims to introduce
users to the basics of command line interaction. The tool is currently part of
a Master's Thesis project and thus will be part of a user study. To this end, the
tool integrates a logging feature so that the developer may collect insights
about the tool's effectiveness. Logging is on by default and log files are
saved at /tmp/tutor-log.txt on unix based systems. To opt-out of sending your
log file to the developer you may supply the -n or --no-upload-log flag when
you start the program.

example:
cli-tutor -n or cli-tutor --no-upload-log

The log file is essentially a copy of what you see and input during the lesson.

The log file also contains the following information:
- All text outputted and inputted during the lesson
- Timestamps
- The supplied identifier name or id
- The hostname of the device or docker container running the program

For more information regarding commands, use the --help/-h flag or help sub-command.
example:
cli-tutor --help or cli-tutor help

The source code is available at https://gitlab.com/qasimwarraich/cli-tutor
This link can also be printed to the terminal by issuing:
'cli-tutor repo'`)
	},
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Prints a url to the git repository for this tool",
	Long:  `All software has a home. This is cli-tutor's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("https://gitlab.com/qasimwarraich/cli-tutor")
	},
}

func init() {
	rootCmd.Flags().BoolP("no-upload-log", "n", false, "Do not send a copy of the log to the developer")
	rootCmd.Flags().BoolP("no-welcome", "x", false, "Do not show welcome message when entering tutor")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(repoCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
