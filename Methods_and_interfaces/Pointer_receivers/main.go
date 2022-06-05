package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func (v *Vertex) Change() int {
	v.x *= 3
	v.y *= 4
	return v.x + v.y
}

func main() {
	help := Vertex{3, 4}
	fmt.Println(help.Change())
	fmt.Println(help.x, help.y)
	fmt.Println(help.Change())
	fmt.Println(help.x, help.y)
}
