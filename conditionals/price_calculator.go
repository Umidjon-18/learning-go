package main

import (
	"fmt"
	"strings"
)

var items = map[string]float64{
	"Apple":     12.2,
	"Banana":    19.12,
	"Orange":    23.5,
	"Pineapple": 34.2,
}

func calculateItemPrice(code string) (bool, float64) {
	price, found := items[code]
	if !found {
		if strings.HasSuffix(code, "_SALE") {
			price, found := items[strings.Replace(code, "_SALE", "", 1)]
			if found {
				salePrice := price * .9
				fmt.Printf("Original price $%.2f, Sale price $%.2f\n", price, salePrice)
				return true, salePrice
			} else {
				fmt.Println("Item not found")
				return false, 0
			}
		}
	}
	salePrice := price * .9
	fmt.Printf("Original price $%2f, Sale price $%2f\n", price, salePrice)
	return true, salePrice

}
