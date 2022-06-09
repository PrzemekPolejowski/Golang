package main

import "fmt"

func maksInt(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {
	tab := []int{1, 3, 5, -2, -2, 11, -10, 50, 5, -20}
	maks := make([]int, len(tab)+1, len(tab)+1)
	wynik := 0
	for i := len(tab) - 1; i >= 0; i-- {
		maks[i] = maksInt(maks[i+1]+tab[i], tab[i])
		wynik = maksInt(wynik, maks[i])
	}
	fmt.Println(wynik)
}
