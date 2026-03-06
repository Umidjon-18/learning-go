package main

import "fmt"

func main() {
	val := 10

	fmt.Printf("Address of val is %+v\n", &val)

	valPtr := &val
	*valPtr = *valPtr + 9

	fmt.Printf("Val = %+v\n", val)
}
