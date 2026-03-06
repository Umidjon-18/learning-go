package main

import "fmt"

func deferExample() {
	defer fmt.Println("Deferred print")
	fmt.Println("Hello from Uzbekistan")
}

func panicExample(shouldPanic bool) {
	if shouldPanic {
		panic("This function needs to panic")
	}
	fmt.Println("This function executed without panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("This function recovered from panic")
		}
	}()

	panicExample(true)
}

func main() {
	// deferExample()
	recoverable()
}
