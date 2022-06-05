package main

import (
	"fmt"
	"sort"
)

func centeredAvg(tab []int) int {
	sort.Ints(tab)
	help := tab[1 : len(tab)-1]
	suma := 0
	for _, v := range help {
		suma += v
	}
	return suma / len(help)
}

func main() {
	var example4_1 = []int{1, 2, 3, 4, 100}
	fmt.Println(centeredAvg(example4_1)) //output: 3

	var example4_2 = []int{1, 1, 5, 5, 10, 8, 7}
	fmt.Println(centeredAvg(example4_2)) //output: 5

	var example4_3 = []int{-10, -4, -2, -4, -2, 0}
	fmt.Println(centeredAvg(example4_3)) //output: -3

}
