package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTasksCmd)
}

type Task struct {
	Id        int64   `db:"id"`
	Content   string  `db:"content"`
	Completed *string `db:"date_completed"`
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  `list all uncompleted tasks by default, flag may be set to list completed tasks instead. If more than 25 tasks are in the list, page flag can be set`,
	Run: func(cmd *cobra.Command, args []string) {

		rows, err := db.Query("select id,content,date_completed FROM tasks order by date_completed, id")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var t Task
			if err := rows.Scan(&t.Id, &t.Content, &t.Completed); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("[ ] %d: %v \n", t.Id, t.Content)
		}
	},
}
