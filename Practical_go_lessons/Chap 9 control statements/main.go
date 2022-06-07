package main

import (
	"math/rand"
	"time"
)

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	return random
}
func main() {
	// Losowanie liczby	fmt.Println(randomNumber())
	
}
