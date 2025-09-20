package main

import (
	"awesomeProject/repository"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// CreateDb Creates a database if it not exist
func CreateDb() error {

	_, err := os.ReadFile("dbNote.sqlite")
	if err == nil {
		return nil
	}

	_, err = os.Create("dbNote.sqlite")
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", "dbNote.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE if not exists NOTES (id integer primary key autoincrement, name text, description text)")
	if err != nil {
		return err
	}

	return nil
}

// DbInit Database registration
func DbInit() (*repository.NoteRepository, *sql.DB) {
	db, err := sql.Open("sqlite3", "dbNote.sqlite")
	if err != nil {
		panic(err)
	}

	var noteRepository repository.NoteRepository = repository.NoteImplementation{Db: db}

	return &noteRepository, db
}

// DbClose Close database connection
func DbClose(db *sql.DB) {
	db.Close()
}
