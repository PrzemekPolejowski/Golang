package main

import "fmt"

func main() {
	x := 23
	pointer := &x
	fmt.Println(pointer)
	fmt.Println(*pointer)
	*pointer = 100
	fmt.Println(x)
}
