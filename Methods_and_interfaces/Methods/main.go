package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x int
	y int
}

func (v Vertex) pow() float64 {
	return math.Pow(float64(v.x), float64(v.y))
}

func main() {
	x := Vertex{5, 5}
	fmt.Println(x.pow())
}
