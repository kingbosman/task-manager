package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of task-manager",
	Long:  `All software has versions. This is task-manager's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Task-manager v0.9 -- HEAD")
	},
}
