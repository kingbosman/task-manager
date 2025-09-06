package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var page string

func init() {
	rootCmd.AddCommand(listTasksCmd)
	listTasksCmd.Flags().StringVarP(&page, "page", "p", "1", "the page of tasks")
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  `list all uncompleted tasks by default, flag may be set to list completed tasks instead. If more than 25 tasks are in the list, page flag can be set`,
	Run: func(cmd *cobra.Command, args []string) {
		pageN, err := strconv.Atoi(page)
		pageSize := 20
		if err != nil || pageN < 1 {
			log.Fatal("failed to get page")
		}
		offset := (pageN * pageSize) - pageSize

		rows, err := db.Query("select id,content, date_completed FROM tasks order by date_completed, id limit ? offset ?", pageSize, offset)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id uint64
			var content string
			var completed any
			if err := rows.Scan(&id, &content, &completed); err != nil {
				log.Fatal(err)
			}

			if completed != nil {
				completed = "x"
			} else {
				completed = " "
			}

			fmt.Printf("[%s] %d: %v \n", completed, id, content)
		}
	},
}
