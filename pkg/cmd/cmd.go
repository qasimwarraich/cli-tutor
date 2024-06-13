package cmd

import (
	"fmt"
	"log"
	"os"

	"cli-tutor/pkg/logger"
	"cli-tutor/pkg/tui"
	"cli-tutor/pkg/tui/tuihelpers"

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
    Web version is available at https://clitutor.chistole.ch`,

	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogging()
		log.Println(args)
		log.Println(cmd.Flags().GetBool("theme"))
		log.Println(cmd.Flags().GetString("theme"))
        

		nowelcomeflag, err := cmd.Flags().GetBool("no-welcome")
		if err != nil {
			log.Println(err)
		}

		if !nowelcomeflag {
			tuihelpers.ProgramWelcome()
		} else {
			tuihelpers.ClearScreen()
		}

		initialWd, _ := os.Getwd()

		tui.StartUI()

		defer os.Remove(initialWd + "/file.txt")
		defer os.Remove(initialWd + "/.hiddenfile.txt")
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
	Short: "Prints information about the tool",
	Long:  `Prints out information regarding cli-tutor`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			`
cli-tutor is a simple command line tutor application that aims to introduce
users to the basics of command line interaction. A local log file is maintained
at /tmp/tutor-log.txt on unix based systems. 

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

var themeCmd = &cobra.Command{
	Use:   "theme",
	Short: "Set theme",
	Long:  `Set the theme for cli-tutor`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inside themeCmdTheme")
		fmt.Printf("Args in themeCmd=%s",args)

	},
}

func init() {
	rootCmd.Flags().BoolP("no-upload-log", "n", false, "Do not send a copy of the log to the developer")
	rootCmd.Flags().BoolP("no-welcome", "x", false, "Do not show welcome message when entering tutor")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(themeCmd)
	themeCmd.Flags().String("theme", "", "Set theme to use with cli-tutor")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
