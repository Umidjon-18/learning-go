package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, pingCh chan string, pongCh chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-pongCh:
			fmt.Println(msg)
			time.Sleep(time.Second)
			pingCh <- fmt.Sprintf("Ping: %v", time.Now())
		}
	}
}
func pong(ctx context.Context, pingCh chan string, pongCh chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-pingCh:
			fmt.Println(msg)
			time.Sleep(time.Second)
			pongCh <- fmt.Sprintf("Pong: %v", time.Now())
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pingChannel := make(chan string)
	pongChannel := make(chan string)

	go ping(ctx, pingChannel, pongChannel)
	go pong(ctx, pingChannel, pongChannel)

	pingChannel <- fmt.Sprintf("Ping: %v", time.Now())

	<-ctx.Done()

}
