package main

import "fmt"

func oddPosition(x []int) []int {
	var newArray []int
	for i, v := range x {
		if i%2 == 1 {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func main() {
	var example1_1 = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(oddPosition(example1_1)) //output [1 3 5]
	var example1_2 = []int{1, -1, 2, -2}
	fmt.Println(oddPosition(example1_2)) //output [-1 -2]

}
