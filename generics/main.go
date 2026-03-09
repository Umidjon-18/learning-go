package main

import (
	"cmp"
	"fmt"
)

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

func Min[T cmp.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Min("abc", "ab"))
}
