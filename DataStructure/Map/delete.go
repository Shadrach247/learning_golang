// Delete Map Element
// Use the delete(map, key) function to remove a key value pair from a map.

package main

import "fmt"

func main() {
	fruits := map[int] string {1:"Apple", 2:"Mango", 3:"Pineapple", 4:"Orange", 5:"Watermelon" }
delete(fruits, 2)
delete(fruits, 4)

fmt.Println(fruits)
}