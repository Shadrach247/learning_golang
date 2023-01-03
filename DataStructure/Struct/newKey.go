// A struct can be declared using the new keyword and assign values using dot notation.
package main

import "fmt"

type Student struct {
	name     string
	subject  string
	class    string
	grade    int
}

func main() {

	var s2     = new(Student)
	s2.name    = "Revelations"
	s2.subject = "C.R.K"
	s2.class   = "2C"
	s2.grade   = 89

	fmt.Println(s2)
}