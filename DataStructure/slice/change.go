// We can use the same index number to change the value of a slice element.

package main
import "fmt"

func main() {
  fruits := []string{"PawPaw", "Orange", "Mango"}

  //change the element of index 2
  fruits[2] = "Apple"

  fmt.Println(fruits)
}