// Program to add elements to a slice

package main

import "fmt"

func main() {
	primeNumbers := []int{2, 3}

// add elements 5, 7 to the slice
primeNumbers = append(primeNumbers, 5, 7)
fmt.Println("Prime Numbers:", primeNumbers)

// We can also add all elements of one slice to another slice using the append() function.
// // Program to add elements of one slice to another
// create two slices
evenNumbers := []int{2, 4}
oddNumbers := []int{1, 3}  

// add elements of oddNumbers to evenNumbers
evenNumbers = append(evenNumbers, oddNumbers...)
fmt.Println("Numbers:", evenNumbers)
}