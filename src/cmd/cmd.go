package cmd

import (
	"fmt"
	"log"
	"os"

	"cli-tutor/src/logger"
	"cli-tutor/src/tui"

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
		tui.StartUI()

		sendlog, err := cmd.Flags().GetBool("no-upload-log")
		if err != nil {
			log.Panicln(err)
		}
		if !sendlog {
			defer logger.UploadLog()
		}
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
