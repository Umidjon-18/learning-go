package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func Sum[T Number](numbers ...T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
}
