// The continue keyword is used to skip an iteration and then continue with the next iteration. 
// Taking the previous example of printing 1 through 5, use the continue keyword to skip printing 4.

package main

import "fmt"

func main() {
for num := 1; num <= 5; num++ {
	if num == 4 {
		continue //continue next iteration
	}
	fmt.Println(num)
	} 
}