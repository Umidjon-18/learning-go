package main

import (
	"fmt"
	"time"
)

func runSwitch() {
	day := "Monday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend! No work")
	case "Monday", "Thuesday":
		fmt.Println("Workdays, a lot of meetings")
	default:
		fmt.Println("Mid-week, not too much work")
	}

	switch hour := time.Now().Hour(); {
	case 6 < hour && hour < 12:
		fmt.Println("Good Morning")
	case hour < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i any) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer %v\n", v)
		case bool:
			fmt.Printf("Boolean %v\n", v)
		case string:
			fmt.Printf("String %v\n", v)
		default:
			fmt.Printf("Unknown type %v\n", v)

		}
	}

	checkType(0)
	checkType("A")
	checkType(false)
}
