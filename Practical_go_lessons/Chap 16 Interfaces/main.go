package main

import "fmt"

type Fish struct {
	length int
	weight int
}
type Snake struct {
	length int
	wiegth int
}

func (f Fish) RetLen() int {
	return f.length
}
func (s Snake) RetLen() int {
	return s.length
}

type Dlugosc interface {
	RetLen() int
}

func main() {
	dorsz := Fish{
		12,
		110,
	}
	mamba := Fish{
		20,
		50,
	}
	help := []Dlugosc{dorsz, mamba}
	for _, v := range help {
		fmt.Println(v.RetLen())
	}
	//fmt.Println("dsd")
}
