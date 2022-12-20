// program to create a slice and print its elements
package main
import (
	"fmt"
	"reflect"
)

func main() {
	// declare slice variable of type integer
	numbers := []int{2, 4, 6, 8, 15}
	var nums = []int{3, 2, 6, 7, 9,}

	// There are 3 ways to find the type of variables in Golang.
	// print the slice
	fmt.Println("Numbers: ", numbers)
	fmt.Println("nums: ", nums)
	fmt.Println(reflect.TypeOf(numbers))
	fmt.Printf("numbers = %T\n", numbers)
	fmt. Println("Numbers = ", reflect.ValueOf(numbers).Kind())	
}