package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tempFile, err := os.CreateTemp("", "hello.c")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("Removing temporary file at: ", tempFile.Name())
		os.Remove(tempFile.Name())
	}()

	_, err = tempFile.Write([]byte("Hello from Go"))
	if err != nil {
		log.Fatal(err)
	}
	tempFile.Close()

	tempDir, err := os.MkdirTemp("", "")
	fmt.Println(tempDir)
	os.Remove(tempDir)
}
