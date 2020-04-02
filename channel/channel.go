package main

import (
  "fmt"
  "time"
)

func delay(c chan int) {
  time.Sleep(5000 * time.Millisecond)
  fmt.Println("done")
  c <- 1
}

func main() {
  c := make(chan int)
  fmt.Println("start")
  go delay(c)
  fmt.Println("end")
  <- c //wait for data from channel
}
