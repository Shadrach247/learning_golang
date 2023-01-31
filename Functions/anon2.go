// Anonymous Function with Parameters

// Program to pass arguments in an anonymous function

package main
import "fmt"

func main() {

  // anonymous function with arguments
  var sum = func(num1, num2 int) {
    sum := num1 + num2
    fmt.Println("Sum is:", sum)
  } 

  // function call
  sum(5, 3)

}