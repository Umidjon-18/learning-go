package main

import "fmt"

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

func (e MathError) Error() string {
	return e.Message
}

func Sum(numbers ...int) int {
	defer fmt.Println("Sum function completed")
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func Division(a, b int) (int, error) {
	if b == 0 {
		err := MathError{
			Operation: "Division",
			InputA:    a,
			InputB:    b,
			Message:   "Dividing to 0 is not allowed",
		}
		return 0, err
	}
	return a / b, nil
}

func main() {
	fmt.Println(Sum(12, 23))

	r1, err1 := Division(13, 2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(r1)
	}

	r2, err2 := Division(13, 0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(r2)
	}

}
