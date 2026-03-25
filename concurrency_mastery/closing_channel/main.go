package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	go func(wg *sync.WaitGroup) {
		for {
			fmt.Println("Loop is running again")
			result, ok := <-jobs
			if ok {
				fmt.Println("Got this message: ", result)
				wg.Done()
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}(&wg)

	for i := range 5 {
		wg.Add(1)
		jobs <- i
		fmt.Println("Sending message: ", i)
	}
	close(jobs)

	wg.Wait()
}

func doubleChannelConcept() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			fmt.Println("Loop is running again")
			result, ok := <-jobs
			if ok {
				fmt.Println("Got this message: ", result)
			} else {
				fmt.Println("Channel closed")
				done <- true
				return
			}
		}
	}()

	for i := range 5 {
		jobs <- i
		fmt.Println("Sending message: ", i)
	}
	close(jobs)

	<-done
}
