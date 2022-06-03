package main

import (
	"fmt"
)

func main() {
	// var table [10]int
	// for i := 0; i < 10; i++ {
	// 	table[i] = int(math.Pow(float64(i), 2))
	// }
	// fmt.Println(len(table))
	// fmt.Println(table)
	// var help []int = table[4:]
	// help = help[:2]
	// fmt.Println(len(help))
	// fmt.Println(help)
	// help[0] = 70
	// fmt.Println(table)
	// tab := []int{2, 3}
	// fmt.Println(tab)
	// var s []int
	// if s == nil {
	// 	fmt.Println("dsd")
	// }
	tab := []int{1, 2, 3}
	fmt.Println(tab)
	slice := tab[1:2]
	fmt.Println(slice)
	table := make([]int, 5, 10)
	fmt.Println(cap(table))
	//Slices of slices
	tablica := []int{1, 2, 3, 4}
	fmt.Println(tablica)
	slice = tablica[1:4]
	fmt.Println(slice)
	slice = slice[1:]
	fmt.Println(slice)
	//Funkcja append
	tablica = append(tablica, 102, 302)
	fmt.Println(tablica)
}
