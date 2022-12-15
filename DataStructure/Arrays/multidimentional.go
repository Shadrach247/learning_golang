// Program to illustrate multidimensional array

package main
import "fmt"

func main() {

  // create a 2 dimensional array
  arrayInteger := [2][2]int{{1, 2}, {3, 4}}

  // access the values of 2d array
  for i := 0; i < 2; i++ {
    for j := 0; j < 2; j++ {
      fmt.Println(arrayInteger[i][j])
    }
  }
}
// Here, we have created a multidimensional array arrayInteger with 2 rows and 2 columns.
// A multidimensional array can be treated like a matrix like this

