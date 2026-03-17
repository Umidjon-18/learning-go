package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DownloadFile(url, dir string) error {
	filename := filepath.Base(url)
	filePath := filepath.Join(dir, filename)

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("File create error: %+v", err)
	}
	defer out.Close()

	fmt.Println("Download started")
	start := time.Now()

	res, err := http.Get(url)
	if err != nil {
		os.Remove(filePath)
		return fmt.Errorf("Http request error: %+v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		os.Remove(filePath)
		return fmt.Errorf("Bad status: %+v", err)
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return fmt.Errorf("File copy error: %+v", err)
	}
	fmt.Printf("Download took %+v\n", time.Since(start))
	return nil
}

func DownloadMultipleFiles(urls []string, dir string) error {
	start := time.Now()
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	for _, url := range urls {
		err := DownloadFile(url, dir)
		if err != nil {
			fmt.Printf("Error while downloading %+v. Error: %+v\n", url, err)
			continue
		}
	}
	fmt.Printf("Sequential download took %+v\n", time.Since(start))
	return nil
}

func ConcurrentDownloader(urls []string, dir string, maxConcurrency int64) error {
	start := time.Now()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	var wg sync.WaitGroup
	limiter := make(chan struct{}, maxConcurrency)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			limiter <- struct{}{}
			defer func() { <-limiter }()
			DownloadFile(url, dir)
		}(url)
	}

	wg.Wait()
	fmt.Printf("Concurrent download took %+v\n", time.Since(start))
	return nil
}

func main() {
	images := []string{
		"https://static.vecteezy.com/system/resources/thumbnails/057/068/323/small/single-fresh-red-strawberry-on-table-green-background-food-fruit-sweet-macro-juicy-plant-image-photo.jpg",
		"https://media.istockphoto.com/id/1550071750/photo/green-tea-tree-leaves-camellia-sinensis-in-organic-farm-sunlight-fresh-young-tender-bud.jpg?s=612x612&w=0&k=20&c=RC_xD5DY5qPH_hpqeOY1g1pM6bJgGJSssWYjVIvvoLw=",
		"https://thumbs.dreamstime.com/b/beautiful-rain-forest-ang-ka-nature-trail-doi-inthanon-national-park-thailand-36703721.jpg",
		"https://cdn.pixabay.com/photo/2016/11/21/06/53/beautiful-natural-image-1844362_640.jpg",
	}
	DownloadMultipleFiles(images, "../images1")
	ConcurrentDownloader(images, "../images2", 3)

}
