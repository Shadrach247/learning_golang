// In computer programming, a recursive function calls itself.
// The recurse() function includes the function call within its body.
// Hence, it is a Go recursive function and this technique is called recursion.

package main
import "fmt"

func countDown(number int) {

// display the number
  fmt.Println(number)

  if number > 0 {
	fmt.Println(number)
// recursive call by decreasing number
	countDown(number - 1)
  } else {
	// ends the recursive function
	fmt.Println("Countdown Stops")
  }

}

func main() {
  countDown(5)
}