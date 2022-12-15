/*
In Go, the switch statement allows us to execute one code block among many alternatives.
The expression after the switch keyword is evaluated. If the result of the expression is matched
the code block will be executed.
In case there is no match, the default code block is executed.
*/

package main

import "fmt"

func main() {
	dayOfWeek := 5
  
	switch dayOfWeek {
  
	  case 1:
		fmt.Println("Sunday")
		  
	  case 2:
		fmt.Println("Monday")
		  
	  case 3:
		fmt.Println("Tuesday")
		  
	  case 4:
		fmt.Println("Wednesday")
		  
	  case 5:
		fmt.Println("Thursday")
		  
	  case 6:
		fmt.Println("Friday")
		  
	  case 7:
		fmt.Println("Saturday")
		  
	  default:
		fmt.Println("Invalid day")

/*
In the above example, I have assigned 5 to the dayOfWeek variable. 
Now, the variable is compared with the value of each case statement.
Since the value matches with case 5, the statement fmt.Println("Thursday") inside the case is executed.
Unlike other programming languages, we don't need to use break after every case. 
This is because in Go, the switch statement terminates after the first matching case.
*/
	}	  
}


// Go switch case with fallthrough
// Program to print the day of Birth using fallthrough in switch
func dob() {

	birthMonth := 2

	switch birthMonth {

	case 1:
	  fmt.Println("January")
		  
	case 2:
	  fmt.Println("February")
		  
	case 3:
	  fmt.Println("March")
	  fallthrough
		  
	case 4:
	  fmt.Println("April")
			 
	case 5:
	  fmt.Println("May")
		  
	case 6:
	  fmt.Println("June")
		  
	case 7:
	  fmt.Println("July")
		  
	default:
	  fmt.Println("Invalid day")

/*
If we need to execute other cases after the matching case, 
we can use fallthrough inside the case statement.
*/
	}
}

func dow() {
// 	Go switch with multiple cases
dayOfWeek := "Sunday"

  switch dayOfWeek {
    case "Saturday", "Sunday":
      fmt.Println("Weekend")

    case "Monday","Tuesday","Wednesday","Thursday","Friday":
      fmt.Println("Weekday")

    default:
      fmt.Println("Invalid day")
// We can also use multiple values inside a single case block. 
// In such case, the case block is executed if the expression matches with one of the case values.
    }
}

/*
In the above example, we have used multiple values for each case:
# case "Saturday", "Sunday" - executes if dayOfWeek is either Saturday or Sunday
# case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday" - executes if dayOfWeek is either one of the value
*/

// Golang switch without expression
func nod() {
numberOfDays := 28

// Switch without any expression
// Program to check if it's February or not using switch without expression
switch {

  case 28 == numberOfDays:
	fmt.Println("It's February")
		   
  default:
	fmt.Println("Not February")
 	}
}
// In the above example, switch doesn't have an expression. 
// Hence, the statement is true.


// Go switch optional statement
/*
In Golang, we can also use an optional statement along with the expression. 
The statement and expression are separated by semicolons
*/

// switch with statement
func d() {
switch day := 6; day {

	case 1:
	fmt.Println("Sunday")

	case 2:
	fmt.Println("Monday")

	case 3:
	fmt.Println("Tuesday")

	case 4:
	fmt.Println("Wednesday")

	case 5:
	fmt.Println("Thursday")

	case 6:
	fmt.Println("Friday") 

	case 7:
	fmt.Println("Saturday")

	default:
	fmt.Println("Invalid Day!")
	}
}
/*
we have used the optional statement day := 6 along with the expression day. 
It matches case 6 and hence, Friday is printed.
*/

