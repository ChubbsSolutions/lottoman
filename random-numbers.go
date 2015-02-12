package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    sum := 1
    for sum < 7 {
    sum ++
    rand.Seed(time.Now().Unix())
    fmt.Println(rand.Intn(48))
    time.Sleep(6328 * time.Millisecond)
  }
}

