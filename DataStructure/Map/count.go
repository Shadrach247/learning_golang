// The len() function returns the total number of elements in a map.

// Get Total Count Copy

package main

import "fmt"

func main() {
fruits := map[int] string {1:"Apple", 2:"Orange", 3:"Watermelon", 4:"Mango", 5:"Pineapple" }

fmt.Println(len(fruits)) //output: 5
}