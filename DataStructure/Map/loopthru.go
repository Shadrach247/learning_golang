// Iterating Map using for Loop
// Similar to arrays, you can iterate through a map using the for loop.

// Map with For Loop

package main
import "fmt"

func main() {
fruits := map[int] string {1:"Pineapple", 2:"Mango", 3:"Orange", 4:"Apple", 5:"Watermelon" }
   
for keyVar, val := range fruits {
	fmt.Printf("%v : %v, ", keyVar, val) // 
	}
}