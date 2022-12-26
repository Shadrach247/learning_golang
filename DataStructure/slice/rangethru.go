// Apart from using for loop, we can also use for range to loop through a slice in Golang

// Program that loops over a slice using for range loop

package main
import "fmt"

func main() {
  numbers := []int{2, 4, 6, 8, 10}

  // for range loop that iterates through a slice
  for _, value := range numbers {
    fmt.Println(value)
  }

}