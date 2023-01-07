// Use shorthand variable syntax to create an instance of the struct

package main

import "fmt"

type Student struct{
	id     int
	name   string
	grade  int
}

func main() {
	s1 := Student{4, "Rukayat", 95}
	fmt.Println(s1)

	s2 := Student{2, "Memunat", 78}
	fmt.Println(s2)
}