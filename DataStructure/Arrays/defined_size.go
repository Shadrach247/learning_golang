/*
An array is a collection of similar types of data. For example,

Suppose we need to record the age of 5 students. 
Instead of creating 5 separate variables, we can simply create an array.
*/


// Program to create an array and prints its elements
package main
import "fmt"

func main() {

	// Syntax
// var array_variable = [size]datatype{elements of array}

   
  // declare array variable of type integer
  // defined size [size=5]
  var arrayOfInteger = [5]int{1, 5, 8, 0, 3}

  fmt.Println(arrayOfInteger)
}