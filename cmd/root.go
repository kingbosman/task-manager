package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var db *sql.DB

var rootCmd = &cobra.Command{
	Use:   "task-manager",
	Short: "task-manager is a simple task manager for keeping track of tasks",
	Long: `A simple yet useable taskmanager. its implicity will make it work keeping track of pending tasks
	the CLI is completely written in golang`,
}

func SetDatabase(conn *sql.DB) {
	db = conn
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
