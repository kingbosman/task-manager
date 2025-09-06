package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/kingbosman/task-manager/cmd"
)

func main() {
	dbPath, err := getDBPath()
	if err != nil {
		fmt.Println("Error getting database path:", err)
		return
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// If file exists, this has already ran once.
	if _, err := os.ReadFile(dbPath); err != nil {
		log.Println("Database not found, starting setup ...")
		if err := createTasksTable(db); err != nil {
			log.Fatal(err)
		}
		log.Println("done with setup")
	}

	cmd.SetDatabase(db)
	cmd.Execute()
}

func createTasksTable(db *sql.DB) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULl,
		date_completed TEXT DEFAULT NULL
	)`); err != nil {
		return err
	}
	return nil

}

func getDBPath() (string, error) {
	// Get the user's home directory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Define the application's data directory. A hidden directory is good practice.
	// .taskmanager is a good name for a directory on macOS/Linux.
	// On Windows, the AppData folder is the standard. This code handles both.
	appDir := filepath.Join(homeDir, ".taskmanager")

	// Create the directory if it doesn't exist.
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}

	// Construct the full path to the database file within the app directory.
	dbPath := filepath.Join(appDir, "taskmanager.db")

	return dbPath, nil
}
