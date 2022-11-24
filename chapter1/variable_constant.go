// Go variable
// A variable is a container that is used to store data. 
// Here's how we can declare a variable in Go programming.

package main

//var xyz int

// xyz - name of the variable
// var - keyword used to declare variable
// int - data type associated with the variable

// Assign Values to Go Variables

// Method 1
//var xyz int = 10
// Method 2
//var xyz = 10
//Method 3
//xyz := 10

import "fmt"

func main() {
 
  // explicitly declare the data type
  var xyz1 int = 10
  fmt.Println(xyz1) // output 10

 // assign a value without declaring the data type
  var xyz2 = 20
  fmt.Println(xyz2) // output 20

  // shorthand notation to define variable
  xyz3 := 30  
  fmt.Println(xyz3) // output 30
  
  
  // Changing Value of a Variable
  // initial value
  xyz4 := 10
  fmt.Println("Initial number value", xyz4) // prints 10

  // change variable value
  xyz4 = 100
  fmt.Println("The changed value", xyz4)  // prints 100

  // Note: In Go, we cannot change the type of variables after it is declared.
  // In the above example, the number variable can only store integer values. 
  // It cannot be used to store other types of data. For example:

  xyz5 := 10
 // assign string data
  xyz5 = "Hello"  // output error code

  // Creating Multiple Variables at Once
  var name, age = "Shadrach", 35
  // Here, "Shadrach" is assigned to the name variable. 
  // Similarly, 35 is assigned to the age variable.

  // The same code above can also be written as:
  name, age := "Shadrach", 35

 // Variables naming rules.
 // A variable name consists of alphabets, digits, and an underscore.
 // Variables cannot have other symbols ( $, @, #, etc).
 // Variable names cannot begin with a number.
 // A variable name cannot be a Go syntax reserved word like int, float, etc.

 // Constant in Go
 // Constants are the fixed values that cannot be changed once declared. 
 // In Go, we use the const keyword to create constant variables. For example,
  const daysOfCode = 100 // initial value
 // Error! Constants cannot be changed
  daysOfCode = 500
 // We cannot use the shorthand notation := to create constants. For example,
 // Error code
  const months := 12


}