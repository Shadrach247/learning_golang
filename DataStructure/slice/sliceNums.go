// In Go programming, we can also create a slice from an existing array.
package main

import "fmt"

func main() {
	numbers := [8]int{11, 25, 30, 47, 55, 60, 78, 82}

// slice the above
	sliceNumbers := numbers[4 : 7]
// 4 is the start of the index and inclusive, while 7 in the end and exclusive.
	
	fmt.Println(sliceNumbers)
}
