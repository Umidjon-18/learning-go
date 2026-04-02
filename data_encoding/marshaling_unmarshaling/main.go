package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Animal struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Color  string `json:"color"`
	Secret Secret `json:"secret"`
}

func (a Animal) String() string {
	return fmt.Sprintf("{Name: %v, Age: %v, Color: %v, Secret: %v}", a.Name, a.Age, a.Color, a.Secret)
}

type Secret struct {
	Something string `json:"something"`
}

func (s Secret) String() string {
	return fmt.Sprintf("{Something: %v}", s.Something)
}

var gotten = `{"name":"Boyka","age":2,"color":"black", "secret":{"something":"Very important"}}`

func main() {
	dog := Animal{
		Name:  "Boyka",
		Age:   2,
		Color: "black",
		Secret: Secret{
			"Very important",
		},
	}
	fmt.Println(dog)

	data, err := json.Marshal(dog)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	fmt.Println("----Returning back----")
	var reversed Animal
	err = json.Unmarshal([]byte(gotten), &reversed)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reversed)

}
