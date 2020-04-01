package main

import (
  "fmt"
  "math"
)

type geometry interface {
  area() float64
  perim() float64
}

type circle struct {
  radius float64
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

type rectangle struct {
  width float64
  height float64
}

func (r rectangle) area() float64 {
  return r.width * r.height
}

func (r rectangle) perim() float64 {
  return 2 * (r.width + r.height)
}

func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}

func main(){
  r := rectangle{width: 2, height: 3}
  c := circle{radius: 3}

  measure(r)
  fmt.Println("-------------------------------")
  measure(c)
}
