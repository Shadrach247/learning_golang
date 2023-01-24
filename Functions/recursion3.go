// Factorial of a number using Go Recursion
package main
import "fmt"

func factorial (num int) int {

  // condition to break recursion
  if num == 0 {
    return 1
  } else {
    // condition for recursion call
    return num * factorial (num-1)
   }
}

func main() {
  var number = 6
  
  // function call
  var result = factorial (number)

  fmt.Println("The factorial of 6 is", result)
}