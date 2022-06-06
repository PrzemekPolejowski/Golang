package main

import "fmt"

type Singers interface {
	Voice() string
}

type Martin struct {
	name    string
	surname string
}

type John struct {
	name    string
	surname string
}

func (v Martin) Voice() string {
	return v.name + v.surname
}

func (v John) Voice() string {
	return v.name + v.surname
}

func main() {
	help := []Singers{Martin{"Martin", "Stan"}, John{"John", "Snow"}}
	for _, v := range help {
		fmt.Println(v.Voice())
	}
	var cos Singers = John{"Somebody", "unknown"}
	fmt.Println(cos.Voice())
}
