/*
Nested For Loop
A for loop can be nested within another for loop. 
For every iteration of the outer loop, the inner loop is executed once.
*/

package main
import ("fmt")

func main() {

    // outer loop
    for i :=0; i < 3; i++ {

      // inner loop 
        for j :=0; j < 4; j++ {
          fmt.Println ("i=",i,"j=", j )
        }

      }

}
/*
In the above nested loop program, the first for loop (outer loop) iterates through 0, 1 & 2 and the inner for loop iterates from 0 (zero) to 3 (three). 
For each iteration of the outer loop, the inner loop is executed once.
*/

