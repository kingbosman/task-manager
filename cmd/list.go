package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTasksCmd)
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  `list all uncompleted tasks by default, flag may be set to list completed tasks instead. If more than 25 tasks are in the list, page flag can be set`,
	Run: func(cmd *cobra.Command, args []string) {

		//TODO: check if null is first, limit 30? also pages
		rows, err := db.Query("select id,content FROM tasks order by date_completed, id")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id uint64
			var content string
			if err := rows.Scan(&id, &content); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("[ ] %d: %v \n", id, content)
		}
	},
}
