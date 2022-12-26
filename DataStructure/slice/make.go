// In Go, we can also create a slice using the make() function.

package main

import (
	"fmt"
	"reflect"
)

func main() {
    
  // create a slice using make()
  nums := make([]int, 8) // When length and capacity is the same.
  numbers := make([]int, 8, 9)  // when length and capacity is different.
    
  // add elements to numbers
  numbers[0] = 4
  numbers[1] = 8
  numbers[2] = 3
  
  // output the length and capacity.
  fmt.Printf("Lenght and Capacity: \tlen%v  \tcap%v\n", len(nums), cap(nums))
  fmt.Println(reflect.ValueOf(nums).Kind())  
  fmt.Println("Numbers:", numbers)
}