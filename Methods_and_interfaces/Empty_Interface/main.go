package main

import "fmt"

func main() {
	var i interface{}
	i = "hello"
	describe(i)

}
func describe(i interface{}) {
	fmt.Printf("%v %T", i, i)
}
