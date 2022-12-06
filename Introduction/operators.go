package main

import "fmt"

func main() {
// Addition, Subtraction and Multiplication Operators
  num1 := 6
  num2 := 2

// + adds two variables
  sum := num1 + num2
  fmt.Printf("%d + %d = %d\n", num1, num2, sum)

// - subtract two variables
  difference := num1 - num2
  fmt.Printf("%d - %d = %d\n",num1, num2,  difference)

// * multiply two variables
  product := num1 * num2
  fmt.Printf("%d * %d is %d\n",num1, num2,  product)

// Golang Division Operator
  num3 := 11
  num4 := 4

// / divide two integer variables
  quotient := num3 / num4
  fmt.Printf(" %d / %d = %d\n", num3, num4, quotient) // output: 11 / 4 = 2

/*
Here, we get the output 2. However, in normal calculation, 11 / 4 gives 2.75. 
This is because when we use the / operator with integer values, 
we get the quotients instead of the actual result.
If we want the actual result we should always use the / operator with floating point numbers.
*/

  num5 := 11.0
  num6 := 4.0

// divide two floating point variables
  result := num5 / num6
  fmt.Printf(" %g / %g = %g\n", num5, num6, result)  //output: 11 / 4 = 2.75

// Modulus Operator in Go
  num7 := 11
  num8 := 4

// % modulo-divides two variables
  remainder := num7 % num8
  fmt.Println(remainder )  //output: 3

/* we get the result 3.
This is because in programming, the modulo operator always returns the remainder after division.
*/

// Increment and Decrement Operator in Go
num := 5

  // increment of num by 1
  num++
  fmt.Println(num)  // 6

  // decrement of num by 1
  num--
  fmt.Println(num)  // 4

// Go Assignment Operators
  newNum := 6
  var example int

// = operator to assign the value of num to result
  example = newNum
  fmt.Println(example)    // 6

// Compound Assignment Operators
  number := 2
  number += 6
// Here, += is additional assignment operator. 
// It first adds 6 to the value of number (2) and assigns the final result (8) to number

/*
Operator	                       Example	   Same as
 += (addition assignment)	       a += b	      a = a + b
 -= (subtraction assignment)	   a -= b	      a = a - b
 *= (multiplication assignment)  a *= b	      a = a * b
 /= (division assignment)	       a /= b	      a = a / b
 %= (modulo assignment)	         a %= b		    a = a % b
*/

// Relational Operators in Golang
 if 5 == 6 {
  print(true)
 } else {
  println(false)
 }
/*
We use the relational operators to compare two values or variables. For example,
Here, == is a relational operator that checks if 5 is equal to 6.

A relational operator returns

true if the comparison between two values is correct
false if the comparison is wrong
*/

/* Logical Operators in Go.
 A logical operator returns either true or false depending upon the conditions.
 && (Logical AND) || (Logical OR) ! (Logical NOT)
*/
}