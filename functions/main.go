package main

import (
	"errors"
	"fmt"
	"strings"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var logger = func(text string) {
	fmt.Println(text)
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func divideName(fullName string) (firstName, lastName string) {
	separated := strings.Split(fullName, " ")
	firstName = separated[0]
	lastName = separated[1]
	return
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger("Hello")
	logger("Hi")
	logger("How are you?")

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())

	fName, lName := divideName("Umidjon Yoqubov")
	fmt.Println(fName, lName)
	fmt.Println(divide(10, 5))
	fmt.Println(divide(3, 0))

}
