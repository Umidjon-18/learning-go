package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func kneader(ctx context.Context, id int, doughChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		batchId := fmt.Sprintf("%d-%d", id, i)
		select {
		case <-ctx.Done():
			return
		case doughChan <- batchId:
			fmt.Printf("Kneader %d: Prepared batch: %v\n", id, batchId)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func baker(ctx context.Context, id int, doughChan <-chan string, rackChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-doughChan:
			fmt.Printf("Batch received: %v. Baker %d\n", v, id)
			// Bake dough for some time
			time.Sleep(50 * time.Millisecond)
			rackChan <- fmt.Sprintf("Bread from Baker %d (Source: %s)", id, v)
		}
	}
}

func main() {
	var doughChan chan string = make(chan string)
	var rackChan chan string = make(chan string)
	var kneaderWg = new(sync.WaitGroup)
	var bakerWg = new(sync.WaitGroup)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	for i := 0; i < 3; i++ {
		kneaderWg.Add(1)
		go kneader(ctx, i, doughChan, kneaderWg)
	}

	for i := 0; i < 3; i++ {
		bakerWg.Add(1)
		go baker(ctx, i, doughChan, rackChan, bakerWg)
	}

	go func() {
		kneaderWg.Wait()
		close(doughChan)
	}()

	go func() {
		bakerWg.Wait()
		close(rackChan)
	}()

	for msg := range rackChan {
		fmt.Println(msg)
	}
	fmt.Println("Factory is closed")
}
