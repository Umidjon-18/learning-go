package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "My name is John"
	fmt.Println(strings.Clone(text))
	fmt.Println(strings.Contains(text, "is"))
	fmt.Println(strings.Compare(text, "is"))
	fmt.Println(strings.HasPrefix(text, "My name"))
	fmt.Println(strings.TrimRight(text, "hn"))
	fmt.Println(strings.TrimSpace("			Hi bob"))
	fmt.Println(strings.ToLower(text))
}
