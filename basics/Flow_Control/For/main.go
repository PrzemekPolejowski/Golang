package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	cos := 1
	for cos < 1000 {
		cos += cos
	}
	fmt.Println(cos)
	// for {
	// 	fmt.Println("jd")
	// }
}
