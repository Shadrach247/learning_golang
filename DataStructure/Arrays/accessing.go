/*
In Go, each element in an array is associated with a number. The number is known as an array index.

We can access elements of an array using the index number (0, 1, 2 …).
*/

// Program to access the array elements
package main
import "fmt"

func main() {
  languages := [5]string{"Go", "Java", "C++", "Python", "Elixir"}

  // access element at index 0
  fmt.Println(languages[0]) // Go

  // access element at index 2
   fmt.Println(languages[2]) // C++

}
/*
The array index always starts with 0. 
Hence, the first element of an array is present at index 0, not 1.
*/