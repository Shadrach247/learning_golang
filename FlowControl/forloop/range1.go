// Range Keyword
// The range keyword is used to iterate through an array, slice, or map.

package main

import "fmt"

func main() {
	colors := [] string{"Blue",  "Yellow",  "Pink"}

	for num, value := range colors {
		fmt.Println (num,"-", value)
	}
}