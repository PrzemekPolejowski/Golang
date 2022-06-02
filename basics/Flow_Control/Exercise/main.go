package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var help int = int(math.Sqrt(x))
	Wieksza := help + 1
	if x-float64(help) < x-float64(Wieksza) {
		return float64(help)
	} else {
		return float64(Wieksza)
	}

}

func main() {
	fmt.Println(Sqrt(16))
}
