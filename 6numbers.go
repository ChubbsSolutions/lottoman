package main

import (
    "fmt"
    "math/rand"
    "time"
)

//Not much to comment. Get 6 random numbers...
func main() {
    sum := 0
    for sum < 6 {
    sum ++
    rand.Seed(time.Now().UnixNano())
    fmt.Println(rand.Intn(35))
  }
}
