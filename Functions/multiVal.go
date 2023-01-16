// In Go, we can also return multiple values from a function.


package main

import "fmt"

// function definition
func calculate(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2

// return two values
 return sum, difference
}

func main() {
	// function call
	sum, difference := calculate(23, 54)
	fmt.Println("Sum:", sum, "Difference:", difference)
}