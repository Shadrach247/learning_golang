// A string is a sequence of characters.
// We can access individual characters of a string
// Like in Arrays, we can use index numbers to access string characters.

package main
import "fmt"

func main() {

	// create and initialize a string
	name := "Nigeria"
	
	// access first character
	fmt.Printf("%c\n", name[0])

	// access fourth character
	fmt.Printf("%c\n", name[3])

	// access last character
	fmt.Printf("%c", name[5])
}