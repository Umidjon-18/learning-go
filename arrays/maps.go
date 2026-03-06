package main

import "fmt"

func runMaps() {
	var m = map[string]bool{}
	fmt.Println(m)
	var n map[string]bool
	fmt.Println(n)
	if n == nil {
		fmt.Println("n is nil")
	}
	n = make(map[string]bool, 0)

	config := make(map[string]int)
	config["key"] = 12
	fmt.Println(config)
}
