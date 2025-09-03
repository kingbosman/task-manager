package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addTaskCmd)
}

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task, the task will be by default uncompleted. You may set a flag to complete it.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: remove ping and add record
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("hello from add")
	},
}
