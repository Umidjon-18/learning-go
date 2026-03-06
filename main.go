package main

import "fmt"

func main() {
	var str string
	str = "Main File"
	println(str)
	interger := 12
	println(interger)

	var dbl = 13.23
	fmt.Printf("%v\n", dbl)

	const PI = 3.14
	fmt.Println(PI)

	println(Sunday)
	println(Monday)
	println(Thuesday)

	var num = 12
	fmt.Println(num)
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)

	printLevelName(Four)

}

type WeekDays int

const (
	Sunday    = 1
	Monday    = 2
	Thuesday  = 3
	Wednesday = 4
	Thursday  = 5
	Friday    = 6
	Saturday  = 7
)

type Numbers int

const (
	One   = 1
	Two   = 2
	Three = 3
	Four  = 4
)

var levelNames = []string{"Bir", "Ikki", "Uch", "To'rt"}

func (l Numbers) String() string {
	if l < One || l > Four {
		return "Unknown"
	}
	return levelNames[l-1]
}

func printLevelName(l Numbers) {
	fmt.Printf("The level name of %d is %s \n", l, l.String())
}
