package main

import (
	"fmt"
	"math/big"
)

func main() {
	//Liczby pierwsze
	tab := []int{1, 2, 3, 6, 7, 11, 13, 17, 19, 100, 200, 1}
	for _, v := range tab {
		if big.NewInt(int64(v)).ProbablyPrime(0) == true {
			fmt.Printf("Liczba %v jest liczbą pierwszą\n", v)
		} else if v == 1 {
			fmt.Println("Jeden jest jedynką")
		} else {
			fmt.Printf("Liczba %v jest liczbą złożoną\n", v)
		}
	}
}
