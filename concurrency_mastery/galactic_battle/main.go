package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func ship(ctx context.Context, id int, ch chan string, rateLimiter chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Ship %d is powering off\n", id)
			return
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Ship %d received: %s\n", id, v)
			select {
			case <-ctx.Done():
				return
			case rateLimiter <- struct{}{}:
				fmt.Printf("Ship %d is firing LASER\n", id)
				time.Sleep(100 * time.Millisecond)
				<-rateLimiter
			}
		}
	}

}

func radar(ctx context.Context, list []chan string) {
	timer := time.NewTicker(100 * time.Millisecond)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			for _, ch := range list {
				select {
				case ch <- "System Alert":
				default:
					// Just skip it
				}
			}
			fmt.Println("--------------------------------")
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	var registry []chan string
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	rateLimiter := make(chan struct{}, 3)

	for i := 0; i < 5; i++ {
		var privateCh chan string = make(chan string)
		registry = append(registry, privateCh)
		id := i
		go ship(ctx, id, privateCh, rateLimiter)
	}
	go radar(ctx, registry)

	<-ctx.Done()

}
