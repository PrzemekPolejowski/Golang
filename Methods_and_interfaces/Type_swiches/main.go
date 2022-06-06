package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println("Nie int, tylko string")
	case bool:
		fmt.Println("Nie int, tylko bool")
	default:
		fmt.Println("To nie jest int")
	}
}

func main() {
	do(12)
	do("strt")
	do(true)
	do(byte(12))
}
