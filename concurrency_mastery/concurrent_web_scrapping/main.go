package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	client http.Client
)

func loadPage(url string, statusCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := client.Get(url)
	if err != nil {
		fmt.Print("Error: ", err)
	} else {
		defer response.Body.Close()
		statusCh <- fmt.Sprintf(url + ": " + response.Status)

	}

}

func main() {
	// set timeout to the client
	client.Timeout = 5 * time.Second

	wg := new(sync.WaitGroup)
	statusCh := make(chan string)

	urls := []string{
		"https://example.com", "https://google.com", "https://github.com", "https://example.com", "https://google.com", "https://github.com", "https://example.com", "https://google.com", "https://github.com",
		"https://example.com", "https://google.com", "https://github.com", "https://example.com", "https://google.com", "https://github.com", "https://example.com", "https://google.com", "https://github.com",
		"https://example.com", "https://google.com", "https://github.com", "https://example.com", "https://google.com", "https://github.com", "https://example.com", "https://google.com", "https://github.com",
		"google.com"}

	for _, url := range urls {
		wg.Add(1)
		go loadPage(url, statusCh, wg)
	}

	go func() {
		wg.Wait()
		close(statusCh)
	}()

	for result := range statusCh {
		fmt.Println(result)
	}
}
