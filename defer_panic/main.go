package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("CustomError{ Code: %d, Message: %s}", c.Code, c.Message)
}

func getError() error {
	return CustomError{Code: 500, Message: "Internal server error"}
}

func errorWrapping() error {
	err := getError()
	return fmt.Errorf("%w. It seems something is broken in serverside.", err)
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover block is running because of panic")
		}
	}()

	panic("Booom")
}

func main() {
	// fmt.Println(getError())
	// fmt.Println(errorWrapping())

	var section = make([][2]int, 10)
	for i := range section {
		if i == 0 {
			section[i][0] = 0
		} else {
			section[i][0] = section[i-1][1] + 1
		}

		if i < 9 {
			section[i][1] = section[i][0] + 100
		} else {
			section[i][1] = 1000 - 1
		}
	}

	fmt.Println("Section after shaped: ", section)
}
