package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, i)
		time.Sleep(time.Second)
		results <- i
		fmt.Printf("Worker %d finished job %d\n", id, i)
	}
}

func main() {
	numberOfJobs := 5

	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)

	for i := range 3 {
		go worker(i, jobs, results)
	}

	for i := range numberOfJobs {
		jobs <- i
	}
	close(jobs)
	for range numberOfJobs {
		<-results
	}

}
