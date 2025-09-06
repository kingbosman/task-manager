package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeTask)
}

var completeTask = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  `Complete a task by the specified id`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Invalid arguments expecting 1 got %d", len(args))
		}
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			log.Fatal("Failed to converse to int")
		}

		_, err = db.Exec(`UPDATE tasks SET date_completed = datetime('now') where id = ?`, id)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Completed task %d", id)
	},
}
