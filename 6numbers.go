package main

import (
    "fmt"
    "math/rand"
    "time"
)

//Generate six random numbers.
func main() {
    i := 0
    var num []int
    for i < 6 {
        rand.Seed(time.Now().UnixNano())
        temp:=rand.Intn(35)
        if IsItUnique(num,temp) {
             fmt.Print(temp,"\n")
             num=append(num,temp)
             i++
        }
    }

}
//Function to make sure the number is unique.
func IsItUnique(slice []int, proposed int) bool {
    for _, ele := range slice {
        if ele == proposed {
            return false
        }
    }
    return true
}
