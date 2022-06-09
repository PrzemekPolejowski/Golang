package main

import (
	"fmt"
	"sort"
)

func main() {
	tab := []int{168, 373, 273, 390, 230, 444, 163, 10, 481, 8, 298, 304}
	sort.Ints(tab)
	fmt.Println(tab)
	szukana := 273
	start, end := 0, len(tab)-1
	wynik := 0
	for start <= end {
		middle := (start + end) / 2
		if tab[middle] <= szukana {
			start = middle + 1
			wynik = middle
		} else {
			end = middle - 1
		}
	}
	fmt.Println(wynik)
}
