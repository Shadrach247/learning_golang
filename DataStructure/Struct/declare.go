// Declare and use a struct

package main
import "fmt"

func main() {

  // declare a struct
  type Person struct {
    name string
    age  int
  }

  // assign value to struct while creating an instance
  person1 := Person{ "Genesis", 15}
  fmt.Println(person1)

  // define an instance
  var person2 Person
  // assign value to struct variables
  person2 = Person {
    name: "Corinthians",
    age: 30,
  }
  fmt.Println(person2)

}