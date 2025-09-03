package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/kingbosman/task-manager/cmd"
)

func main() {
	db, err := sql.Open("sqlite3", "./taskmanager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected to sqllite3")
	// TODO: create table if not exist
	cmd.SetDatabase(db)
	cmd.Execute()
}
