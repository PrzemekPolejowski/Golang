package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 3
	// switch x {
	// case 1:
	// 	fmt.Println("Wrong")
	// case 2:
	// 	fmt.Println("right")
	// default:
	// 	fmt.Println("JD")
	// }
	x := time.Now()
	fmt.Println(x.Hour())
	switch {
	case x.Hour() < 12:
		fmt.Println("Good Morning")
	case x.Hour() < 20:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}
}
