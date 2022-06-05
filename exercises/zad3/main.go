package main

import (
	"fmt"
	"strings"
)

func listOfDigits(x string) []string {
	return strings.Split(x, "")
}

func main() {
	var example3_1 = "123"
	fmt.Println(listOfDigits(example3_1)) //output [1 2 3]
	var example3_2 = "400"
	fmt.Println(listOfDigits(example3_2)) //output [4 0 0]
}
