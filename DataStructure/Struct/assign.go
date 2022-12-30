// You can assign values to the fields of a struct using curly bracket and assign it to a variable. 
// This is called instance. You can create multiple instances of the struct with different values.
package main
import "fmt"

type Student struct{
	id       int
	class    string
	name     string
	subject  string
	grade    int
}
func main() {
//you must assign values to all the fields in the same order as in the struct declaration.
	var s5 = Student{6, "3B", "Obi", "Biology", 75}
	fmt.Println(s5)

//Or use named parameters like "name:" to assign values in a different order or to some fields.
	s6 := Student{name:"Oluwa", subject: "Maths", grade: 90, class: "2D"}
	fmt.Println(s6)

}