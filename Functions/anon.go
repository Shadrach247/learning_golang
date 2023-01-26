// In Go, we can create a function without the function name, known as an anonymous function.
// What to do, is assign the anonymous function to a variable and then use the variable name to call the function.

// Go Anonymous Function
package main
import "fmt"

func main() {

  // anonymous function
  var favorite = func() {
    fmt.Println("Titanic, is the name of my favorite movie?")
  }

  // function call
  favorite()

}