package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/kingbosman/task-manager/cmd"
)

func main() {
	db, err := sql.Open("sqlite3", "./taskmanager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// If file exists, this has already ran once.
	if _, err := os.ReadFile("./taskmanager.db"); err != nil {
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
