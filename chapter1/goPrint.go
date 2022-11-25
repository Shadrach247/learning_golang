/* We use these three functions to print output messages in Go programming.

 fmt.Print()
 fmt.Println()
 fmt.Printf()

Note: All these functions are defined under the fmt package.
*/

package main

// import fmt package
import "fmt"

func main() {
// Go fmt.Print()
  fmt.Print("Hello, ")
  fmt.Print("World!")
  name := "Shadrach"
  fmt.Print("Name: ", name)

// fmt.Println() prints a new line at the end by default.
  currentSalary := 50000
  // Go fmt.Println
  fmt.Println("Hello")
  fmt.Println("World!")
  fmt.Println("Current Salary:", currentSalary)

// The fmt.Printf(). The 'f' is a format specifier.
  currentAge := 35
  // Go fmt.Printf()
  fmt.Printf("Age = %d", currentAge)

/* In Go, every data type has a unique format specifier.
  Data Type	Format Specifier
  integer	%d
  float	%g
  string	%s
  bool	%t
*/

  // Using %g to print Float Values
  var annualSalary float64 = 65000.5
  fmt.Printf("Annual Salary: %g", annualSalary)

  // Using format specifiers to hold value of a variable
  var myName = "John"
  age := 35
  fmt.Printf("%s is %d years old.", myName, age)

}