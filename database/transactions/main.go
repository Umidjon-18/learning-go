package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"learning-go/database/transactions/repository"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var userTableSchema = `CREATE TABLE IF NOT EXISTS users (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
hashed_password BLOB NOT NULL,
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

var profileTableSchema = `CREATE TABLE IF NOT EXISTS profiles (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_id INTEGER NOT NULL,
bio TEXT,
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

var insertUserQuery = `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`

var getUserByEmailQuery = `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`

var insertProfileQuery = `INSERT INTO profiles (user_id, bio) VALUES (?, ?)`

var getProfileByUserIdQuery = `SELECT id, user_id, bio, created_at FROM profiles WHERE user_id = ?`

func main() {
	dbName := "users.db"
	_ = os.Remove(dbName)
	db, err := connectToDatabase(dbName)
	checkError(err)

	defer func() {
		checkError(db.Close())
	}()

	err = createTable(db, userTableSchema)
	checkError(err)
	err = createTable(db, profileTableSchema)
	checkError(err)

	repository := repository.NewSqlRepository(db)

	_, err = repository.CreateUser("Umidjon 1", "umidjon@gmail.com", "secretPassword")
	_, err = repository.CreateUser("Umidjon 2", "umidjonxomidjonovich@gmail.com", "secretPassword")
	_, err = repository.CreateUser("Umidjon 3", "umidjonyoqubov@gmail.com", "secretPassword")
	checkError(err)
	err = printUsers(repository)
	checkError(err)

}

func printUsers(repository repository.UserRepository) error {
	result, err := repository.GetUsers()
	checkError(err)
	users, err := json.MarshalIndent(result, "", "")
	checkError(err)
	fmt.Println(string(users))
	return nil
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	checkError(err)
	err = db.Ping()
	checkError(err)
	return db, nil
}

func createTable(db *sql.DB, schema string) error {
	_, err := db.Exec(schema)
	return err
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
