// you can use the underscore '_' instead of the variable name

package main
import "fmt"

func main() {
	colors := [] string{"Brown",  "Yellow",  "Purple"}
    
    for _ , value := range colors {
        fmt.Println(value)
    }
}