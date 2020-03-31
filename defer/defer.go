package main

import "fmt"

func main() {
  for i := 0; i < 5; i++{
    defer fmt.Println(i) // stack olarak dizip erteler
  }
  fmt.Println("Done!")
}
