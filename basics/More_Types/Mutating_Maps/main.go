package main

import "fmt"

func main() {
	mapa := make(map[string]int)
	mapa["cos"] = 2
	mapa["ds"] = 100
	delete(mapa, "cos")
	help, check := mapa["cos"]
	fmt.Println(mapa, help, check)
}
