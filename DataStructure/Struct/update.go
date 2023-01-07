// You can update the values of an instance of the struct using dot notation

package main

import "fmt"

type Student struct {
	id int
	name string
	grade int
}

func main() {
s1 := Student{3, "Halima", 9}
fmt.Println(s1)

s1.id = 10
s1.name = "Sophia"
s1.grade = 3

fmt.Println(s1.name, s1.id, s1.grade)
}