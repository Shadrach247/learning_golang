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

/*
LOGICAL OPERATORS IN GO
It returns either 'true' or 'false' depending upon the condition.
Examples.
*/
numb1 := 7
numb2 := 14
numb3 := 7
var ans bool

// returns false because numb1 > numb2 is false
ans = (numb1 > numb2) && (numb1 == numb3)
fmt.Printf("Result of AND operator is %t \n", ans)

// returns true because numb1  == numb3 is true
ans = (numb1 > numb2) || (numb1 == numb3)
fmt.Printf("Result of OR operator is %t \n", ans)

// return false because numb1 == numb3 is true
ans = !(num1 == numb3)
fmt.Printf("Result of NOT operator is %t \n", ans)

/*
Boolean Expression are used to create decision-making programs. Suppose we want to create a
program that determines whether a person can vote or not. We can use a boolean expression
to check if the age of that person is greater than 18. If true, the person can vote. If false,
cannot vote.
*/
}