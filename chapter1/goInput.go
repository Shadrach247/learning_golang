// In Go, we use the scan() function to take input from the user. For example,

package main

import (
  "fmt"
  "log"
)

func main() {

/*  
Takes input value from the user and assigns it to the name variable.
Go programming provides three different variations of the Scan() method:

fmt.Scan()
fmt.Scanln()
fmt.Scanf()
*/

// var name, city string
var name string
var age int
var temperature float32
var sunny bool 

/*
  // takes input value for name
  fmt.Print("Enter your name: \n")
  fmt.Scan(&name)
  fmt.Println("Where do you live?: ")
  fmt.Scan(&city)

  fmt.Printf("Name: %s\nCity: %s\n", name, city)
*/

fmt.Print("Enter your name and age: ")
// Go fmt.Scan
n, err := fmt.Scan(&name, &age)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("scanned %d arguments\n", n)
fmt.Printf("%s is %d years old\n", name, age)

// Go fmt.Scanln
  var newName string
	var newAge int

// Input name from the user
	fmt.Print("Enter name: ")
	fmt.Scanln(&newName)

// Input age from the user
	fmt.Print("Enter age: ")
	fmt.Scanln(&newAge)

// Print the entered values
	fmt.Println("Name: ", newName)
	fmt.Println("Age : ", newAge)

  var num1 int
	var num2 int
	var sum int

// Input numbers
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

// Calculate sum
	sum = num1 + num2
// Printing the values & sum
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)

// Go fmt.Scanf
// take float input
fmt.Println("Enter the current temperature:")
fmt.Scanf("%g", &temperature)
  
// take boolean input
fmt.Println("Is the day sunny?")
fmt.Scanf("%t", &sunny)
  
fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)

/*
Data Type	Format Specifier
integer	        %d
float	        %g
string	        %s
bool	        %t

Each code block will need to be run separately to get the expected output.
*/

}