package main

import "fmt"

func write(x string) {
	fmt.Println(x)
}
func main() {
	// defer fmt.Println("Hello")
	// defer fmt.Println("World")
	// defer fmt.Println("Test")
	fmt.Println("Hello")
	for i := 0; i < 11; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("World")
}
