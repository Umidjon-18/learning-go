package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fpath := filepath.Base("Umidjon/Desktop/main.app")
	fmt.Println(fpath)
	fmt.Println(filepath.Join("config", "assets.json"))
	fmt.Println(filepath.Clean("Umidjon/../Desktop/main/home/./dir/../../go.app"))
	fmt.Println(filepath.Ext("Umidjon/../Desktop/main/home/./dir/../../go.app"))
}
