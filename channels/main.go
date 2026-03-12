package main

import (
	"fmt"
	"time"
)

func testUnbuffered() {
	fmt.Println("UNBUFFERED CHANNEL")
	channel := make(chan int)
	start := time.Now()
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Sending: ", i)
			channel <- i
			fmt.Println("Sent: ", i)
		}
	}()

	for range 5 {
		time.Sleep(500 * time.Millisecond)
		msg := <-channel
		fmt.Println("Received: ", msg)
	}
	fmt.Println("Time took: ", time.Since(start))
}

func testBuffered() {
	fmt.Println("BUFFERED CHANNEL")

	channel := make(chan int, 5)
	start := time.Now()
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Sending: ", i)
			channel <- i
			fmt.Println("Sent: ", i)
		}
	}()

	for range 5 {
		time.Sleep(500 * time.Millisecond)
		msg := <-channel
		fmt.Println("Received: ", msg)
	}
	fmt.Println("Time taken: ", time.Since(start))
}

func main() {
	testUnbuffered()
	testBuffered()
}
