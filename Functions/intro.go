/*
A function is a group of statements that exist within a program for the purpose of performing a specific task. 
At a high level, a function takes an input and returns an output.
Function allows you to extract commonly used block of code into a single component.
The single most popular Go function is main(), which is used in every independent Go program.
*/


package main

import "fmt"

// define a function
func greet() {
	fmt.Println("Good Morning")
}

// addition function
func addNumbers() {
	no1 := 12
	no2 := 8

	sum := no1 + no2
	fmt.Println("Sum:", sum)
}

func main() {

	// function call
	greet()
	addNumbers()
}



