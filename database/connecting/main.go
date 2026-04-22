package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
hashed_password BLOB NOT NULL,
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

func main() {
	dbName := "demo.db"

	// _ = os.Remove(dbName)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("Closing db")
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}
