package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  radius float64
}

type Rectangle struct {
  width, height float64
}

func (r Rectangle) area() float64 {
  return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
  return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
  return 2 * math.Pi * c.radius
}

func measure(s Shape) {
  fmt.Println(s)
  fmt.Println(s.area())
  fmt.Println(s.perimeter())
}

func main() {
  r := Rectangle{width: 3, height: 4}
  c := Circle{radius: 5}

  measure(r)
  measure(c)
}
