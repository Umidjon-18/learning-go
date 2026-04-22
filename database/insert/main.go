package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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
	dbName := "users.db"
	db, err := createDb(dbName)
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
	err = createTable(db, schema)
	if err != nil {
		log.Fatal(err)
	}
	id, err := addUser(db, "Umidjon", "umidjonyoqubov@gmail.com", "secretPassword")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last inserted id: ", id)

}

func createDb(dbName string) (*sql.DB, error) {
	fmt.Println("Creating db: ", dbName)
	conn, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createTable(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func addUser(db *sql.DB, name, email, password string) (*int64, error) {
	query := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	res, err := db.Exec(query, name, email, hp)
	if err != nil {
		return nil, err
	}
	lsId, err := res.LastInsertId()

	return &lsId, nil
}
