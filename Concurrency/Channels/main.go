package main

import "fmt"

func sum(x []int, c chan int) {
	suma := 0
	for _, v := range x {
		suma += v
		fmt.Println(v)
	}
	c <- suma
}

func main() {
	tab := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int)
	go sum(tab[(len(tab)/2):], c)
	go sum(tab[:(len(tab)/2)], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
