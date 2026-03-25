package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// func ping(cancelCh <-chan struct{}, pingCh chan<- struct{}, pongCh <-chan struct{}, logCh chan<- string) {
// 	for {
// 		select {
// 		case <-cancelCh:
// 			return
// 		case <-pongCh:
// 			time.Sleep(time.Second)
// 			select {
// 			case pingCh <- struct{}{}:
// 				select {
// 				case logCh <- "Ping":
// 				default:
// 				}
// 			case <-cancelCh:
// 				logCh <- "Finished"
// 				return
// 			}
// 		}
// 	}
// }
// func pong(cancelCh <-chan struct{}, pongCh chan<- struct{}, pingCh <-chan struct{}, logCh chan<- string) {
// 	for {
// 		select {
// 		case <-cancelCh:
// 			return
// 		case <-pingCh:
// 			time.Sleep(time.Second)
// 			select {
// 			case pongCh <- struct{}{}:
// 				select {
// 				case logCh <- "Pong":
// 				default:
// 				}
// 			case <-cancelCh:
// 				logCh <- "Finished"
// 				return

// 			}
// 		}
// 	}
// }

// func logger(logCh <-chan string) {
// 	for log := range logCh {
// 		fmt.Println(log)
// 	}
// }

// func main() {

// 	pingCh := make(chan struct{})
// 	pongCh := make(chan struct{})
// 	cancelCh := make(chan struct{})
// 	logCh := make(chan string, 20)

// 	go ping(cancelCh, pingCh, pongCh, logCh)
// 	go pong(cancelCh, pongCh, pingCh, logCh)
// 	go logger(logCh)

// 	pingCh <- struct{}{}
// 	logCh <- "Ping"
// 	time.Sleep(5 * time.Second)
// 	close(cancelCh)
// 	time.Sleep(time.Second)
// 	close(logCh)
// }

var (
	result string
)

func addToFinal(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel
	result += letter
	wg.Done()
	fmt.Println("Added: ", result)
}

func capitalize(letterChannel chan string, letter string, wg *sync.WaitGroup) {
	upper := strings.ToUpper(letter)
	wg.Done()
	letterChannel <- upper
}

func main() {
	runtime.GOMAXPROCS(2)
	wg := new(sync.WaitGroup)
	var pipe chan string = make(chan string)

	inputString := "hello my friends, how are you"

	input := []byte(inputString)

	for i := 0; i < len(input); i++ {
		fmt.Println("For loop index: ", i)
		wg.Add(2)
		go capitalize(pipe, string(input[i]), wg)
		go addToFinal(pipe, wg)
		wg.Wait()

	}

	fmt.Println(result)

}
