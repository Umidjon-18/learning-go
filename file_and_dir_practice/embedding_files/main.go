package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var helloContent string

//go:embed texts
var byeContent embed.FS

func main() {
	fmt.Println(helloContent)
	content, err := byeContent.ReadDir("texts")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}
