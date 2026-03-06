package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers[0:5])
	fmt.Println(numbers[5:])

	numbers = append(numbers, 10, 10, 10, 14)
	fmt.Println(numbers)
	numbers = slices.Insert(numbers, 12, 17)
	fmt.Println(numbers)

}
