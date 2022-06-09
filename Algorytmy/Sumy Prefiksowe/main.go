package main

import "fmt"

func sumaPrzedzial(x int, y int, tab []int) int {
	return tab[y+1] - tab[x]
}

func main() {
	tab := []int{1, 2, 3, 56, 10, -2}
	prefiks := [7]int{}
	for i := 1; i < len(prefiks); i++ {
		prefiks[i] = prefiks[i-1] + tab[i-1]
	}
	fmt.Println(prefiks)
	//Suma z danego przedziału
	fmt.Println(sumaPrzedzial(3, 5, prefiks[:]))
}
