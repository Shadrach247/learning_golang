package main

import "fmt"

// We use data types in Golang to determine the type of data associated with variables. 
// For example,
var age int
// 'int' is a data type that specifies that the 'age' variable can store integer data.

// The basic data types in Golang: 

// int	Integer numbers. | 8014, 0, -1, 8632
// float	Numbers with decimal points. |  50.3, 320.74635, -89.20
// complex	Complex numbers.  |  2+4i, -9.5+18.3i
// string	Sequence of characters.	|  "Hello Golang!", "5 is less than 7"
// bool	either true or false. |  true, false
// byte	A byte (8 bits) of non-negative integers. |  9, 216, 77
// rune	Used for characters. Internally used as 32-bit integers. 'a', '7', '<'

// 1. Integer Data Type
// Integers are whole numbers that can have both zero, positive and negative values but no decimal values.
// You can declare multiple variables at once in the same line.

var id, number int

// In Go programming, there are two types of integers:

// signed integer int - can hold both positive and negative integers
// unsigned integer uint - can only hold positive integers

// There are different variations of integers in Go programming.
// Data type	Size
// int/uint	either 32 bits (4 bytes) or 64 bits (8 bytes)
// int8/uint8	8 bits (1 byte)
// int16/uint16	16 bits (2 bytes)
// int32/uint32	32 bits (4 bytes)
// int64/uint64	64 bits ( 8 bytes)

func main() {
  var xyz6 int
  var xyz7 int

  xyz6 = 5
  xyz7 = 10

  fmt.Println(xyz6)   // output 5
  fmt.Print(xyz7)     // output 10

// 2. Float Data Type
// The float type is used to hold values with decimal points. e.g, 6.7, -34.2

// Keywords used: float32, float64
// An example,
// var salaryA float64
// There are two sizes of floating-point data in Go programming.
// Data Type	Size
// float32	32 bits (4 bytes)
// float64	64 bits (8 bytes)
// Note: If we define float variables without specifying size explicitly, the size of the variable will be 64 bits.
// the size of the variable is 64
// salaryB := 5676.3

// Program to illustrate float32 and float64 with example
  var salaryA float32
  var salaryB float64

  salaryA = 1234.5678910

  // can store decimals with greater precision
  salaryB = 1234.5678910

  fmt.Println(salaryA)  // output  101234.5679
  fmt.Println(salaryB)  // output  1234.567891


  // 3. String Data Type
  // A string is a sequence of characters. For example, "Hello friend", "Hi Dear"

  // Keyword: string
  var message string
  // In Go, we use either double quotes or backticks to create strings.

  message = "Hello World "
  var post string =  `Hello World`

  fmt.Println(message)
  fmt.Println(post)

  // Boolean Data Type
  // The boolean data type has one of two possible values either true or false.
 // Keyword: bool
  var myBoolValue bool
  myBoolValue = false

  fmt.Println(myBoolValue)  // false
}