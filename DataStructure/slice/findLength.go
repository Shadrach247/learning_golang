// We use the len() function to find the number of elements present inside the slice. 

// Program to find the length of a slice
package main
import "fmt"

func main() {

  // create a slice of numbers
  numbers := []int{1, 8, 4, 9, 2, 5, 7}

  // find the length of the slice
  length := len(numbers)

  fmt.Println("Length:", length)

}