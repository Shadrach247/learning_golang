/*
package main


import "fmt"

func main() {

// Go fmt.Scanln()
var myName string
var myAge int

// take name and age input
fmt.Print("Enter your name?: ")
fmt.Scanln(&myName, &myAge)
// fmt.Print("How old are you?: ")

fmt.Printf("Name: %s\nAge: %d\n", myName, myAge)

}

import (
    "fmt"
)

func main() {

  // the message we try here would be:
  // James is a fantastic footballer and 23 years old
  // or something follwoing the same format

  var name string
  var number int
  var temp string

  // taking input and storing in variable using the buffer string
  fmt.Println("What do you have to say?: ")
  fmt.Scanln(&name, &temp, &number)

  // print out new string using the extracted values
  fmt.Printf ("%s is a fantastic %s and %d years old", name, temp, number );
}
*/

package main

import (
	"fmt"
)

func main() {
	num := 6
  var result int

  // = operator to assign the value of num to result
  result = num
  fmt.Println(result)    // 6
  
}