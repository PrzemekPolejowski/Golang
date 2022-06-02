package main

import (
	"fmt"
	"math"
)

func main() {
	var table [10]int
	for i := 0; i < 10; i++ {
		table[i] = int(math.Pow(float64(i), 2))
	}
	fmt.Println(len(table))
	fmt.Println(table)
	var help []int = table[4:]
	help = help[:2]
	fmt.Println(len(help))
	fmt.Println(help)
	help[0] = 70
	fmt.Println(table)
	tab := []int{2, 3}
	fmt.Println(tab)
	var s []int
	if s == nil {
		fmt.Println("dsd")
	}
}
