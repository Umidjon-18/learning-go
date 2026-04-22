package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

func main() {
	dbName := "users.db"
	_ = os.Remove(dbName)
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
	id, err := addUser(db, "Umidjon", "umidjonxomidjonovich@gmail.com", "secretPassword")
	if err != nil {
		log.Fatal(err)
	}

	id, err = addUser(db, "Umidjon", "umidjonyoqubov@gmail.com", "secretPassword")
	if err != nil {
		log.Fatal(err)
	}

	id, err = addUser(db, "Umidjon", "umidjon@gmail.com", "secretPassword")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last inserted id: ", id)

	res, err := getUserByEmail(db, "umidjonyoqubov@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	userJs, err := json.MarshalIndent(res, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(userJs))

	result, err := getAllUsers(db)
	if err != nil {
		log.Fatal("Error", err)
	}
	for _, user := range result {
		u, err := json.MarshalIndent(user, "", "")
		if err != nil {
			log.Fatal("Error", err)
		}
		fmt.Println(string(u))
	}

}

func createDb(dbName string) (*sql.DB, error) {
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

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email=?`

	row := db.QueryRow(query, email)
	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email, hashed_password, created_at FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
