// Add or Update Elements
// You can add a new element or update an existing one in the map using the map[key]. 
// The following updates the value of an existing key 2, and adds a new key 6, and its value to the map.

package main

import "fmt"

func main() {
	fruits := map[int] string {1:"Orange", 2:"Mango", 3:"WaterMelon"}
fruits[2] = "Banana" //update
fruits[6] = "Pineapple" //add
fruits[4] = "Apple" //add

fmt.Println(fruits) 
}