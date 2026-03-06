package main

import "fmt"

func runIfElse() {
	num := 30

	if num > 30 {
		fmt.Println("Greater than 30")
	} else {
		fmt.Println("Less or equal to 30")
	}

	permissions := map[string]bool{
		"Umidjon": true,
		"Doniyor": false,
	}

	if value, exists := permissions["Umidjon"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key was not found")
	}
}
