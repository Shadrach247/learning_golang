// The map[key] returns two values, first the value itself, and second, 
// true or false depending upon if the specified key exists or not. 
// It returns true if a key exists; otherwise returns false.

package main

import "fmt"

func main() {
	colleagues := map[int] string {9: "Shadrach", 2: "Ibrahim", 7: "Benjamin", 5: "Paul"}

// check for an existing key
value1, exist1 := colleagues[9]
fmt.Println(value1, exist1)

// check an non-existing key
value2, exist2 := colleagues[6]
fmt.Println(value2, exist2)

}

// In the above example, key 9 exists in the emplist map and so it returns the value associated with key 9 and returns true. 
// For key value 6, it returns the default value of string which is a blank string "", and also returns false because key 6 does not exist.