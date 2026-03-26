package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID         int
	Complexity int // Represents time in ms
}

func collector(ctx context.Context, taskChan chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	id := 1
	timer := time.NewTicker(100 * time.Millisecond)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			taskChan <- Task{
				ID:         id,
				Complexity: 100 * id,
			}
			id++
		}
	}
}

func worker(ctx context.Context, id int, tasks <-chan Task, results chan<- int, gpu chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is shutting down\n", id)
			return
		case t, ok := <-tasks:
			if !ok {
				return
			}
			fmt.Printf("Processing task %d\n", t.ID)
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d is shutting down\n", id)
				return
			case gpu <- struct{}{}:
				results <- t.ID
				time.Sleep(time.Duration(t.Complexity))
				<-gpu
			}

		}
	}
}

func aggregator(results <-chan int, finalWg *sync.WaitGroup) {
	defer finalWg.Done()
	for r := range results {
		fmt.Printf("Task %d is done\n", r)
	}

}

func main() {
	// Setup channels, context, and waitgroups
	var tasks chan Task = make(chan Task, 10)
	var results chan int = make(chan int, 10)
	var gpu chan struct{} = make(chan struct{}, 2)
	var producerWg = new(sync.WaitGroup)
	var workerWg = new(sync.WaitGroup)
	var aggregatorWg = new(sync.WaitGroup)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// Launch goroutines
	producerWg.Add(1)
	go collector(ctx, tasks, producerWg)

	go func() {
		producerWg.Wait()
		close(tasks)
	}()

	for i := 0; i < 5; i++ {
		workerWg.Add(1)
		id := i
		go worker(ctx, id, tasks, results, gpu, workerWg)
	}

	go func() {
		workerWg.Wait()
		close(results)
	}()
	aggregatorWg.Add(1)
	go aggregator(results, aggregatorWg)

	// Handle graceful exit
	aggregatorWg.Wait()

}
