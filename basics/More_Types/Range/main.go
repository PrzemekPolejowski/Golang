package main

import "fmt"

func main() {
	var tab = []int{1, 2, 3, 4, 5, 1000}
	for _, i := range tab {
		fmt.Println(i)
	}
}
