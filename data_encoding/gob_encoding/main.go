package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
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

	var buf bytes.Buffer
	gobEncoder := gob.NewEncoder(&buf)
	err := gobEncoder.Encode(user1)
	if err != nil {
		log.Fatal(err)
	}

	gobDecoder := gob.NewDecoder(&buf)
	var user2 User
	gobDecoder.Decode(&user2)
	fmt.Println(user2)
}
