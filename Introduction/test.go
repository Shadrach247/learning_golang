package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

  /*
  var num1 int = 90
  var num2 float32 = 8.7
  var sum float32
 
  // addition of different data types
  sum = float32(num1) + num2

  fmt.Printf("Sum is %g",sum)
  */
}