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
  fmt.Print("Hello,")
  fmt.Print("World!")
  name := "Shadrach"
  fmt.Print("Name: ", name) // outputs adds no new line. Hello,World!Name: Shadrach

// fmt.Println() prints a new line at the end by default.
  currentSalary := 700000
  // Go fmt.Println
  fmt.Println("Hello")   // 'ln' adds new line to each output. Hello
  fmt.Println("World!")  // World
  fmt.Println("Current Salary:", currentSalary) // Current Salary: 700000

// The fmt.Printf(). The 'f' is a format specifier.
  currentAge := 36
  // Go fmt.Printf()
  fmt.Printf("Age = %d\n", currentAge) // Age = 36.

/* In Go, every data type has a unique format specifier.
  Data Type	Format Specifier
  integer	%d
  float	%g
  string	%s
  bool	%t
*/

  // Using %g to print Float Values
  var annualSalary float64 = 34000000
  fmt.Printf("Annual Salary: %g\n", annualSalary) // Annual Salary: 34000000.0

  // Using format specifiers to hold value of a variable
  var myName = "Shadrach"
  age := 36
  fmt.Printf("%s is %d years old.\n", myName, age) // Shadrach is 36 years old

}