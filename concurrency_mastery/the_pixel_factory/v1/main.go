package main

import (
	"context"
	"fmt"
	"time"
)

type Task struct {
	ID         int
	Complexity int // in milliseconds
}

func producer(ctx context.Context, taskCh chan Task) {
	id := 1
	timer := time.NewTicker(100 * time.Millisecond)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			task := Task{
				ID:         id,
				Complexity: id * 10,
			}
			select {
			case <-ctx.Done():
				return
			case taskCh <- task:
			}
			id++
		}
	}
}

func dispatcher(ctx context.Context, taskCh <-chan Task, poolCh chan<- Task) {
	for {
		select {
		case <-ctx.Done():
			return
		case t, ok := <-taskCh:
			if !ok {
				return
			}
			fmt.Printf("Dispatcher received task %v\n", t)
			select {
			case <-ctx.Done():
				return
			case poolCh <- t:
				fmt.Printf("Dispatcher sent task %v to worker pool\n", t)
			}
		}
	}
}

func worker(id int, ctx context.Context, poolCh <-chan Task) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is shutting down\n", id)
			return
		case t, ok := <-poolCh:
			if !ok {
				return
			}
			time.Sleep(time.Duration(t.Complexity))
			fmt.Printf("Worker %d processed task: %v\n", id, t)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var taskCh chan Task = make(chan Task)
	var poolCh chan Task = make(chan Task, 2)
	go producer(ctx, taskCh)
	go dispatcher(ctx, taskCh, poolCh)
	for i := 0; i < 4; i++ {
		id := i
		go worker(id, ctx, poolCh)
	}

	<-ctx.Done()
}
