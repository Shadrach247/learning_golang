// In Go, each element in a slice is associated with a number. 
// The number is known as a slice index

// Program to access the slice elements
package main

import "fmt"

func main() {

  languages := []string{"Go", "JavaScript", "Elixir", "C", "Dartlang", "Rustlang"}

  // access element at index 0
  fmt.Println(languages[1])

  // access element at index 2
  fmt.Println(languages[4])

  // access elements from index 0 to index 2
  fmt.Println(languages[0:3]) 

}
// The slice index always starts with 0. 
// The first element of a slice is present at index 0, not 1.