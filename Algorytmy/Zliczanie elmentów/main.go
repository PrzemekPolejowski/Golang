package main

import "fmt"

func main() {
	tab := []int{1, 2, 4, 4, 4, 4, 4, 2, 6, 7, 8, 3, 1, 5, 6, 3, 5, 6, 3, 2, 3, 4, 4, 4, 8, 7, 9, 9, 9}
	licznik := [10]int{}
	for _, v := range tab {
		licznik[v]++
	}
	for i := 1; i < 10; i++ {
		fmt.Printf("Liczba %v występuje %v razy\n", i, licznik[i])
	}

}
