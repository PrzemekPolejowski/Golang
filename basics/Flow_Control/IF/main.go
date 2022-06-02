package main

import "fmt"

func main() {
	x := 10
	if x < 5 {
		fmt.Println("Tak")
	} else if x < 10 {
		fmt.Println("Mniejsza od 10")
	} else {
		fmt.Println("Wieksza lub ronwa 10")
	}

}
