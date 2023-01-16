// In Go, we can declare variables in two different scopes: local scope and global scope.
// A variable scope specifies the region where we can access a variable.

// Program to illustrate local variables

package main

import "fmt"

func addNumbers() {
	// local variables
	var sum int
	sum = 12 + 8
}

func main() {

	addNumbers()

	// cannot access sum out of its local scope
	fmt.Println("Sum is", sum)
}