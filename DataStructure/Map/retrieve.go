// The values in a map can be retrieved by passing the key in the square bracket, e.g. map[key].

package main

import "fmt"

func main() {
	colleagues := map[int] string {9: "Shadrach", 2: "Ibrahim", 7: "Benjamin", 5: "Paul"}

	fmt.Println(colleagues[7])
	fmt.Println(colleagues[9])
	fmt.Println(colleagues[5])
	fmt.Println(colleagues[3])
	fmt.Println(colleagues[2])

}