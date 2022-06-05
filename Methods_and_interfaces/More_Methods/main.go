package main

import "fmt"

type Test float64

func (v Test) Abs() float64 {
	if v < 0 {
		return float64(v) * -1
	}
	return float64(v)
}

func main() {
	help := Test(-3)
	fmt.Println(help.Abs())
}
