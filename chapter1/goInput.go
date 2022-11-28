// In Go, we use the scan() function to take input from the user. For example,

package main

import "fmt"

func main() {
  var name string
  var city string

  // takes input value for name
  fmt.Print("Enter your name: ")
  fmt.Scan(&name, &city)

  fmt.Printf("Name: %s\nCity: %s\n", name, city)

/*  
fmt.Scan(&name)
Takes input value from the user and assigns it to the name variable.
Go programming provides three different variations of the Scan() method:

fmt.Scan()
fmt.Scanln()
fmt.Scanf()
*/

  // Go fmt.Scanln()
var myName string
var myAge int
    
// take name and age input
fmt.Println("Enter your name and age:")
fmt.Scanln(&myName, &myAge) 
fmt.Printf("Name: %s\nAge: %d\n", myName, myAge)

var newName string
var newAge int

// take name and age input using format specifier
  fmt.Println("Enter your name and age:")
  fmt.Scanf("%s %d", &newName, &newAge) 
  fmt.Printf("Name: %s\nAge: %d\n", newName, newAge)


  var temperature float32
	var sunny bool
  
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
*/
}