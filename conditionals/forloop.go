package main

func main() {
	var items = []string{"Apple", "Banana", "Chocolate", "Cola"}

	for i := 0; i < len(items); i++ {
		println(items[i])
	}

	var target = 10

	for target > 0 {

		if target%2 == 0 {
			println(target)
		}
		target--
	}

	for index, value := range items {
		println(index, value)
	}

	runIfElse()
	runSwitch()
	calculateItemPrice("Apple")
	calculateItemPrice("Banana_SALE")
	calculateItemPrice("Bread")
}
