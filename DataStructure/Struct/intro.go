/* A struct (short for "structure") is a collection of data fields with declared data types.
 Golang has the ability to declare and create own data types by combining one or more types,
 including both built-in and user-defined types. Each data field in a struct is declared with
 a known type, which could be a built-in type or another user-defined type.
 Structs are the only way to create concrete user-defined types in Golang. Structs types are
 declared by composing a fixed set of unique fields. Structs can also be considered as a template
 for creating a data record, like an employee record or an e-commerce product.
 The declaration starts with the keyword type, then a name for the new struct, and finally the
 keyword struct. Within the curly brackets, a series of data fields are specified with a name
 and a type.
*/
package main

import "fmt"

type identifier struct{
	field1 string   // < data_type
	field2 int      // < data_type
	field3 float64  // < data_type
}
func main() {
	fmt.Println(identifier{"Green", 1, 2.0 })
}