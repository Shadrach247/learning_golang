// Program to access the individual elements of struct

package main

import "fmt"


type Student struct {
	id int
	name string
	grade int
}

func main() {
s1 := Student{5, "Genesis", 80}

//Accessing s1 values
fmt.Println("Name: ", s1.name) 
fmt.Println("ID: ", s1.id)  
fmt.Println("Grade: ", s1.grade)  
}