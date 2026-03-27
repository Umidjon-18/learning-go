package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
	once    sync.Once
)

func something(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go something(wg)
	}

	wg.Wait()

	for i := 0; i < 500; i++ {
		go once.Do(func() {
			fmt.Println("Do something")
		})
	}

	fmt.Println(counter)

}
