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
	Long:  `Add a new task, the task will be by default uncompleted.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Invalid arguments expecting 1 got %d", len(args))
		}
		res, err := db.Exec(`INSERT INTO tasks(content) VALUES(?)`, args[0])
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
