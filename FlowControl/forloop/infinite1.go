/*
A loop that never terminates is called an infinite loop.

If the test condition of a loop never evaluates to true, the loop continues forever. For example,
*/

// Program to infinitely iterate through the loop
package main

import "fmt"

func main() {

  for i := 0; i <= 10; i-- {
    fmt.Println("Welcome to Programiz") // infinite
  }
}