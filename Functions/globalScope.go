/* A global variable refers to a variable that is defined outside a function. 
Global variables can be accessed throughout the program or within any function in the defined package.
Typically, global variables are defined on top of the program before the main function. 
After declaration, a global variable can be accessed and changed at any part of the program.
*/

package main

import "fmt"

var x = 10  // a global variable
 
func main() {
	
    fmt.Println(x)
    foo()
	editable()
}
 
func foo() {
    fmt.Println(x * 2)
}

// Global variables in Go are editable. 
// Any code in any of the functions can edit the global variable.
func editable() {
	x = 20 // editing global variable
	fmt.Println(x)   
}

// GoLang allows variables of the same name if the scopes are different. 
// So, we can define a variable named ‘x’ inside the function too. 
// In this case, the local variable will get more preference than the global variable.
