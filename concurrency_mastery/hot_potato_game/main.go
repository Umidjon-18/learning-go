package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	expireTime  = 24
	playerCount = 10
)

func player(id int, ch chan int, quitCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-quitCh:
			return
		case v := <-ch:
			{
				if v >= expireTime {
					fmt.Printf("Booom. Current count: %d. My id: %d\n", v, id)
					close(quitCh)
				} else {
					fmt.Printf("Current count: %d. My id: %d\n", v, id)
					time.Sleep(50 * time.Millisecond)
					ch <- v + 1
				}
			}
		}
	}
}

func main() {
	wg := new(sync.WaitGroup)
	var pipe chan int = make(chan int, 1)
	var quit chan struct{} = make(chan struct{})
	wg.Add(playerCount)
	for i := range playerCount {
		go player(i, pipe, quit, wg)
	}
	pipe <- 0
	wg.Wait()
	fmt.Println("Game over. ")
}
