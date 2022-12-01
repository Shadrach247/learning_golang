package main

import "fmt"

// type casting is the process of converting the value of one data type (integer, float, string) to another data type.
func main() {
	var floatValue float32 = 5.45

// Manually performing the type conversion, is called explicit type casting in Go.
// type conversion from float to int
	var intValue int = int(floatValue)
   
	fmt.Printf("Float Value is %g\n", floatValue)
	fmt.Printf("Integer Value is %d", intValue)

/*
Implicit Type Casting in Go
In implicit type casting, one data type is automatically converted to another data type. 
However, Go doesn't support implicit type casting.
*/

// initialize integer variable to a floating-point number
   var number int = 4.34 // error, Go doesn't support implicit type conversion.

   fmt.Printf("Number is %g", number)

// Add int and float number using Go type casting
   var number1 int = 20
   var number2 float32 = 5.7
   var sum float32
 
  // addition of different data types
   sum = float32(number1) + number2

   fmt.Printf("Sum is %g",sum)
}
