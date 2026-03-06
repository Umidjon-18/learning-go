package main

import (
	"errors"
	"fmt"
	"time"
)

var DivisionByZeroError = errors.New("Divided by zero")
var TooLargeError = errors.New("Number is too large")

type OpError struct {
	Op      string
	Message string
	Code    int
	Time    time.Time
}

func (e OpError) Error() string {
	return e.Message
}

func CreateNewError(op, message string, code int, t time.Time) error {
	err := OpError{Op: op, Message: message, Code: 100, Time: t}
	return err
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	} else if a > 1000 {
		return 0, TooLargeError
	}
	return a / b, nil

}

func main() {

	result, err := divide(1001, 1)

	if errors.Is(err, DivisionByZeroError) {
		fmt.Println(err)
	} else if errors.Is(err, TooLargeError) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	var newErr = CreateNewError("Do something else", "Do something please, error happened", 100, time.Now())
	fmt.Println(newErr)
}
