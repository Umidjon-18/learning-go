package main

import (
	"encoding/json"
	"os"
)

type User struct {
	ID          int            `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Permissions map[string]int `json:"permissions"`
}

func main() {
	user1 := User{
		ID:          1,
		Username:    "John",
		Email:       "john@mail.com",
		Permissions: map[string]int{"one": 1, "two": 2},
	}
	jdr := json.NewEncoder(os.Stdout)
	jdr.Encode(user1)

}
