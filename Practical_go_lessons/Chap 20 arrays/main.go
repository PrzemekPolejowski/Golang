package main

import "fmt"

func main() {
	//var tab []int
	//Two dimensional arrays
	var tab [3][2]int
	tab[1][1] = 5
	fmt.Println(tab, cap(tab))
}
