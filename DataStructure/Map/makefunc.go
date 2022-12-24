// The built-in make() function can be used to create an empty map by specifying the key and value type

// create an empty map
package main
import "fmt"

func main() {
var fruits = make(map[int]string)
fmt.Println(fruits) //output: map[]

// Add fruits to the map
fruits[1] = "Apple"
fruits[2] = "Orange"
fruits[3] = "Watermelon"
fmt.Println(fruits)
// delete Apple
delete(fruits, 1)
fmt.Println(fruits)  // map[2:Orange 3:Watermelon]
// update orange with pineapple instead
fruits[2] = "Pineapple"
fmt.Println(fruits)
// Add mongo & Banana to the list of fruits
fruits[5] = "Mango"
fruits[6] = "Banana"
fmt.Println(fruits)
}