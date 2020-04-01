package main

import (
  "fmt"
  "time"
)

func sum(num0, num1 float64) {
  fmt.Println("Calculating...")
  time.Sleep(2000 * time.Millisecond)
  fmt.Println(num0 + num1)
}

func main() {
  fmt.Println("start")
  go sum(11, 33)
  time.Sleep(4000 * time.Millisecond) //total run time is 4 second
  fmt.Println("end")
}
