package cmd

import (
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
		//TODO:
		// do input
		input := "change me now"
		res, err := db.Exec(`INSERT INTO tasks(content) VALUES(?)`, input)
		if err != nil {
			log.Fatal(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("added task with id: %d", id)
	},
}
