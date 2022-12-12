/*
Golang Blank Identifier in for loop.
In Golang, we have to use every variable that we declare inside the body of the for loop. If not used, it throws an error. We use a blank identifier _ to avoid this error.

Let's understand it with the following scenario.

  numbers := [5] int {11, 22, 33, 44, 55}

  // index variable is declared but not used
  for index, item := range numbers {
    // throws an error
    fmt.Println(item)
  }

}
Here, we get an error message index declared but not used.

To avoid this, we put _ in the first place to indicate that we don't need the first component of the array( index ). Let's correct the above example with a blank identifier.
*/

// Program to print the elements of the array
package main
import "fmt"
 
func main() {
 
  numbers := [5] int {11, 22, 33, 44, 55}

  for _, item := range numbers {
    fmt.Println(item)
  }

}
