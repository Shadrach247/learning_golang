// To change an array element, we can simply reassign a new value to the specific index

// Program to change element by assigning the desired index with a new value

package main
import "fmt"

func main() {
  weather := [4]string{"Rainy", "Sunny", "Cloudy", "Snowy"}

  // change the element of index 1
  weather[1] = "Funny"

  fmt.Println(weather)
}