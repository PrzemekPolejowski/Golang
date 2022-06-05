package main

import "fmt"

func cumulativeSum(array []int) []int {
	//fmt.Println(array)
	for i := range array {
		if i > 0 {
			array[i] += array[i-1]
			//	fmt.Print(v, " ")
		}
	}
	return array
}

func main() {
	var example2_1 = []int{1, 1, 1}
	fmt.Println(cumulativeSum(example2_1)) //output [1 2 3]
	var example2_2 = []int{1, -1, 3}
	fmt.Println(cumulativeSum(example2_2)) // output [1 0 3]

}
