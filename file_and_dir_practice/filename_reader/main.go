package main

import (
	"fmt"
	"math/rand"
	"os"
)

var (
	names = []string{"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
		"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
		"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
		"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
		"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
		"sparkling", "throbbing", "shy", "wandering", "withered", "wild", "black",
		"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
		"restless", "divine", "polished", "ancient", "purple", "lively", "nameless"}
	exts = []string{
		"json",
		"txt",
		"go",
		"py",
		"csv",
	}
)

func generateDummyFiles(count int) error {
	err := os.MkdirAll("test_files", 0755)
	if err != nil {
		return fmt.Errorf("Cannot create test_files folder: %w", err)
	}
	useds := make(map[string]struct{})
	namesLength := len(names)
	extsLength := len(exts)
	for range count {
		randomName := names[rand.Intn(namesLength)]
		randomExt := exts[rand.Intn(extsLength)]
		name := fmt.Sprintf("%v.%v", randomName, randomExt)
		if _, ok := useds[name]; ok {
			continue
		}
		useds[name] = struct{}{}
		f, err := os.Create(fmt.Sprintf("test_files/%v", name))
		if err != nil {
			return fmt.Errorf("Error while creating file: %w", err)
		}
		f.Close()
	}

	return nil
}

func readFileNames() ([]os.DirEntry, error) {

	d, err := os.ReadDir("test_files")
	if err != nil {
		return nil, err
	}

	return d, nil
}

func main() {
	err := generateDummyFiles(10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(readFileNames())
}
