package main

import "fmt"

func main() {
	cos := 3
	var p *int
	p = &cos
	fmt.Println(p, *p)
	var point *int
	fmt.Println(point)
}
