package  main

import "fmt"

func main() {
    //using var
    var colleaguesID = map[int] string {10:"Shadrach",  21:"Ibrahim",  32:"Loveth",  43:"Ismael",  54:"Pricillia"}
    fmt.Println(colleaguesID)
 
    //shorthand syntax
    yearsOfExperience := map[string] int {"Shadrach":3, "Ibrahim":5, "Loveth":4, "Ismael":7, "Pricillia":2} 
    fmt.Println(yearsOfExperience)
}