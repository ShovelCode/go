package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Length, Width float64
}

func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

func printArea(s Shape) {
    fmt.Println(s.Area())
}

func main() {
    c := Circle{Radius: 5}
    printArea(c)
    
    r := Rectangle{Length: 3, Width: 4}
    printArea(r)
}
