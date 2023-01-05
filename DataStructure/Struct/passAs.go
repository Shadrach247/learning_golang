// A struct can be passed to a function as an argument in the same way as any other variable.

package main

import "fmt"

type Student struct {
		id int
		name string
		grade int
	}

func main() {
	var s1 Student
	s1.id = 2
	s1.name = "Damilola"
	s1.grade = 75

	//pass s1 to studentDetails()
	studentDetails(s1)
}

func studentDetails(s Student) {
	fmt.Println("Name: ",  s.name) 
	fmt.Println("ID: ",  s.id)  
	fmt.Println("Grade: ",  s.grade)  
}
/*
In the above program, 
the studentDetails() function accepts an input parameter of type Student. 
The variable s1 of type Student is declared in the main function and is passed 
as an argument to the studentDetails() function.
*/
