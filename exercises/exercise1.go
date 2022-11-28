package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.80
	// price converted to float64
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")
	// price converted to float64
	var total float64 = float64(price) + tax 
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	// availableFunds converted to float64
	fmt.Println("Within budget?", total <= float64(availableFunds))
}