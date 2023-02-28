// In Go, we can also pass anonymous functions as arguments to other functions. 
// In that case, we pass the variable that contains the anonymous function.

// Anonymous Function as Arguments to Other Functions

package main
import "fmt"

var sum = 0

// regular function to calculate square of numbers
func findSquare(num int) int {
  square := num * num
  return square
}

func main() {

  // anonymous function that returns sum of numbers
  sum := func(number1 int, number2 int) int {
    return number1 + number2
}

  // function call
  result := findSquare(sum(7, 5))
  fmt.Println("Result is:", result)

}