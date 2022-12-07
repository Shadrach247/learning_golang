/*
 The If statement is used to run a block of code when a certain condition is met.
 true - statements inside the body of If are executed, while
 false - statements inside the body of If are not executed.
*/

package main

import "fmt"

func main() {
	number := 15

	// syntax
	// true if number is greater than 2
	if number > 2 {
		fmt.Printf("%d is a greater than 2\n", number)
	}

	// The If statement may have an optional Else block.
	// Check if number is less than 12
	if number < 12 {
		fmt.Printf("%d is less than 12\n", number)
	} else {
		fmt.Printf("%d is greater than 12\n", number)
	}

	/*
	If...else statement is used to execute a block of code among two alternatives.
	However, if you need to make a choice between more than two alternatives,
	then we use the Else...if statement.
	*/

	// syntax
	if number == 0 {
		// code block 1
	} else if number == 1 {
		// code block 2
	} else {
		// code block 3
	}

	// Go nested if statement
	// You can also use an If statement inside of an If statement.
	
	// An example
	number1 := 9
    number2 := 22

  // outer if statement
  if number1 >= number2 {

  // inner if statement
  if number1 == number2 {
    fmt.Printf("Result: %d == %d", number1, number2)
    // inner else statement
  } else {
    fmt.Printf("Result: %d > %d", number1, number2)
  } 

  // outer else statement
  } else {
    fmt.Printf("Result: %d < %d", number1, number2)
  }

/*
If the outer condition number1 >= number2 is true
inner if condition number1 == number2 is checked
if condition is true, code inside the inner if is executed
if condition is false, code inside the inner else is executed
If the outer condition is false, the outer else statement is executed.
*/


}