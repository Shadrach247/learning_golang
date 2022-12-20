// We can use the copy() function to copy elements of one slice to another.

// Program to copy one slice to another
package main
import "fmt"

func main() {

  // create two slices
  primeNumbers := []int{1, 3, 5, 7}
  numbers := []int{6, 4, 9}

  // copy elements of primeNumbers to numbers
  copy(numbers, primeNumbers)

  // print numbers
  fmt.Println("Numbers:", numbers)
}
/*
Here, only the first 3 elements of primeNumbers are copied to numbers. 
This is because the size of numbers is 3 and it can only hold 3 elements.
The copy() function replaces all the elements of the destination array with the elements.
*/