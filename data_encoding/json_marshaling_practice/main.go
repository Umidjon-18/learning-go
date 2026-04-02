package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID        int             `json:"id"`
	Username  string          `json:"username"`
	Email     string          `json:"email"`
	Password  string          `json:"-"`
	SecureKey json.RawMessage `json:"secure_key"`
}

var marshaledString = `{"id":1,"username":"John","email":"john@mail.com", "secure_key":"Very important key", "unknown_key":"something"}`

func main() {
	user1 := User{
		ID:       1,
		Username: "John",
		Email:    "john@mail.com",
		Password: "johnshiddenpassword",
	}

	marshaledUser, err := json.Marshal(user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshaledUser))
	var result User
	err = json.Unmarshal([]byte(marshaledString), &result)
	if err != nil {
		log.Fatal(err)
	}
	var s string
	json.Unmarshal(result.SecureKey, &s)
	b64Encoded := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(b64Encoded)
	decoded, err := base64.URLEncoding.DecodeString(b64Encoded)
	fmt.Println("Decoded: ", string(decoded))
	var r User
	dec := json.NewDecoder(bytes.NewReader([]byte(marshaledString)))
	dec.DisallowUnknownFields()
	err = dec.Decode(&r)

	// fmt.Println(s)
	// fmt.Println(r)
	// fmt.Println(result)
}
