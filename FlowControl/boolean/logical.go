package main

import "fmt"

/*
LOGICAL OPERATORS IN GO
It returns either 'true' or 'false' depending upon the condition.
Examples.
*/

func main() {
numb1 := 7
numb2 := 14
numb3 := 7
var ans bool

// returns false because numb1 > numb2 is false
ans = (numb1 > numb2) && (numb1 == numb3)
fmt.Printf("Result of AND operator is %t \n", ans)

// returns true because numb1  == numb3 is true
ans = (numb1 > numb2) || (numb1 == numb3)
fmt.Printf("Result of OR operator is %t \n", ans)

// return false because numb1 == numb3 is true
ans = !(numb1 == numb3)
fmt.Printf("Result of NOT operator is %t \n", ans)

/*
Boolean Expression are used to create decision-making programs. Suppose we want to create a
program that determines whether a person can vote or not. We can use a boolean expression
to check if the age of that person is greater than 18. If true, the person can vote. If false,
cannot vote.
*/
}