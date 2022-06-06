package main

import "fmt"

func Compare[T comparable](x []T, y T) int {
	licznik := 0
	for _, v := range x {
		if v == y {
			licznik++
		}
	}
	return licznik
}

func main() {
	test := []int{1, 1, 3, 1}
	test2 := []string{"d", "jd", "jd"}
	fmt.Println(Compare(test, 4))
	fmt.Println(Compare(test2, "jd"))
}
