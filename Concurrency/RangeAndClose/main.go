package main

import "fmt"

func main() {
	c := make(chan int, 5)
	c <- 2
	c <- 5
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}
