// We can also use index numbers to initialize an array.

package main
import "fmt"

func main() {

// declare an array
var arrayOfIntegers[7] int
 
// elements are assigned using index
arrayOfIntegers[0] = 3
arrayOfIntegers[1] = 8
arrayOfIntegers[2] = 15
arrayOfIntegers[3] = 12
arrayOfIntegers[4] = 29
arrayOfIntegers[5] = 17
arrayOfIntegers[6] = 23

fmt.Println(arrayOfIntegers)
}