package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rtm",
	Short: "rtm is a simple task manager for keeping track of tasks",
	Long: `A simple yet useable taskmanager. its implicity will make it work keeping track of pending tasks
	the CLI is completely written in golang`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
