/* In programming, a loop is used to repeat a block of code
In Golang, we use the for loop to repeat a block of code until the specified condition is met.

Here's the syntax of the for loop in Golang.

for initialization; condition; update {
  statement(s)
}

Here,

The initialization initializes and/or declares variables and is executed only once.
Then, the condition is evaluated. If the condition is true, the body of the for loop is executed.
The update updates the value of initialization.
The condition is evaluated again. The process continues until the condition is false.
If the condition is false, the for loop ends.
*/

// Program to print the first 5 natural numbers

package main
import "fmt"

func main() {

  // for loop terminates when i becomes 6
  for i := 1; i <= 5; i++ {
    fmt.Println(i)
  }

}

