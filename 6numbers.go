package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    sum := 1
    for sum < 2 {
    sum ++
    rand.Seed(time.Now().Unix())
    fmt.Println(rand.Intn(35))
    time.Sleep(6328 * time.Millisecond)
  }
}

