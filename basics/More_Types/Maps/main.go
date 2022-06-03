package main

import "fmt"

type Osoba struct {
	imie     string
	nazwsiko string
}

func main() {
	mapa := make(map[string]int)
	mapa["cos"] = 5
	mapa["jd"] = 10
	fmt.Println(mapa)
	help := map[string]Osoba{
		"dsdsd": {"sdsd", "sdds"},
	}
	help["12345678900"] = Osoba{
		"Damian",
		"Czekolada",
	}
	fmt.Println(help)
}
