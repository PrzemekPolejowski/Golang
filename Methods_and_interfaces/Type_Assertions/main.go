package main

import "fmt"

func main() {
	var i interface{} = "hello"
	help, ok := i.(string)
	fmt.Println(help, ok)
}
