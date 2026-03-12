package main

import (
	"fmt"
	"sync"
	"time"
)

func saySomething(text string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println(text)
}

func main() {

	var wg sync.WaitGroup

	totalGroup := 5
	for i := range totalGroup {
		wg.Add(1)
		go saySomething(fmt.Sprintf("Hello from go routine %d", i), time.Second, &wg)
	}
	fmt.Println("Hello from main go routine")
	fmt.Println("Bye from main go routine")

	wg.Wait()
}
