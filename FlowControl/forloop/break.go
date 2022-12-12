/*
Break Keyword
The break keyword is used to break the execution of code and exit the for loop.

In the following example, the code stops execution and exits the loop when num equals 4.
*/

package main

import "fmt"

func main() {
// break Keyword
for num := 1; num <= 5; num++ {
	if num == 4 {
		break //exit for loop from here
	}
	fmt.Println(num)
	}
} 