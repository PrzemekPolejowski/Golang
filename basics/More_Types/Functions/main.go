package main

import "fmt"

func dobule(x int) int {
	return (x * *&x)
}

func main() {
	fmt.Println(dobule(10))
}
