package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("42a")
	fmt.Println(i, err)
}
