// Math and comparison operations in Go require that the included values be of the
// same type. If they're not, you'll get an error when trying to run your code.

// The same is true of assigning new values to variables. if the type of
// value being assigned doesn't match the declared type of the variable, you'll get an error.

// The solution is to use conversions, which let you convert a value from one type to
// another type. You just provide the type you want to convert a value to, immediately
// followed by the value you want to convert in parentheses.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 2
	//conversion from int to float64 and assignment with Golang short variable declaration.
	result := float64(myInt)
	//relect is a package. While TypeOf is a function of the reflect package.
	fmt.Println(reflect.TypeOf(result))
}