package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    x, y float64
}

func (v Vertex) abs() float64 {
    return math.Sqrt(v.x*v.x + v.y*v.y)
}

func absFunc(v Vertex) float64 {
    return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
    v := Vertex{3 ,4}
    fmt.Println(v.abs())
    fmt.Println(absFunc(v))

    p := &Vertex{4, 3}
    fmt.Println(p.abs())
    fmt.Println(absFunc(*p))
}
