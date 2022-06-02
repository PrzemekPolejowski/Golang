package main

import (
	"fmt"
)

type Table struct {
	x int
	y int
}

func main() {
	// v := Table{1, 2}
	// v.x = 50
	// fmt.Println(v.x)
	// pointer := &v
	// pointer.y = 150
	// fmt.Println(v)
	v1 := Table{1, 2}
	v2 := Table{x: 2}
	v3 := Table{}
	pointer := &Table{50, 20}
	fmt.Println(v1, pointer, v2, v3)
}
