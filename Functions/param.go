// program to illustrate function parameters

package main

import "fmt"

// A function with 2 parameters
func addNumbers(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("Sum", sum)
}

func main() {

	// pass parameters in function call
	addNumbers(33, 8)
}