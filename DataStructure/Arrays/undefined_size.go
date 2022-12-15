// An array can also be declared without specifying its size

package main
import "fmt"

func main() {
   
  // declare array variable of type string
  // undefined size
  var arrayOfString = [...]string{"Hello", "Programiz"}

  fmt.Println(arrayOfString)
}
/*
Note: Here, if [] is left empty, it becomes a slice. 
So [...] is a must if we want an undefined size array.
*/