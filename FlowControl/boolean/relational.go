/* 
Go Booleans (Relational and Logical Operators)
A boolean data type represents logical entities. 
It can have two possible values: true or false.
We use the 'bool' keyword to create boolean-type variables.
*/

package main

import "fmt"


func main() {

	var yes bool = true
    var no bool = false

    fmt.Println("The boolean values are", yes, "and", no)

// Relational (comparison) Operators
number1 := 7
number2 := 2

result := number1 > number2    // a boolean expression.
fmt.Println("Result:", result)
// Relational Operators use boolean values(true and false) to return the validity.

// Examples
num1 := 14
num2 := 19
var output bool

// equal to operator
output = (num1 == num2)
fmt.Printf("%d == %d returns %t \n", num1, num2, output)

// not equal to operator
output = (num1 != num2)
fmt.Printf("%d != %d returns %t \n", num1, num2, output)

// greater than operator
output = (num1 > num2)
fmt.Printf("%d > %d returns %t \n", num1, num2, output)

// less than operator
output = (num1 < num2)
fmt.Printf("%d < %d returns %t \n", num1, num2, output)
}