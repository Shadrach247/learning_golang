// In Go, we can use range with for loop to iterate over an array. For example,

package main

import "fmt"
 
func main() {
 
  // create an array
  numbers := [5] int {11, 22, 33, 44, 55}

  // for loop with range
  for item := range numbers {
    fmt.Println(numbers[item])
  }

}

