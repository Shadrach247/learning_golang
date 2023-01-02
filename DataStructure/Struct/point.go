// An intance of a struct can be created using pointer symbol '&'.

package main

import "fmt"

type Student struct {
	name     string
	subject  string
	class    string
	grade    int
}

func main() {

	var s1 = &Student{"Genesis", "Maths", "1A", 92}

	fmt.Println(s1)
}