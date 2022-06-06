package main

import (
	"fmt"
	"math"
)

type Vertx struct {
	x int
	y int
}

func (v *Vertx) Scale() {
	v.x *= v.x
	v.y *= v.y
}

func (v *Vertx) Abs() float64 {
	return math.Sqrt(float64(v.x) + float64(v.y))
}
func main() {
	help := Vertx{3, 4}
	fmt.Println(help.Abs())
	help.Scale()
	fmt.Println(help.Abs())
}
