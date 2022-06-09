package main

import (
	"fmt"
	"time"
)

type DataUrodzenia struct {
	dzien   time.Weekday
	miesiac time.Month
	rok     int
}
type Osoba struct {
	Imie     string
	Nazwisko string
	Wzrost   int
	DataUrodzenia
}

func main() {

	Michał := Osoba{
		Imie:          "Michał",
		Nazwisko:      "Dekarz",
		Wzrost:        190,
		DataUrodzenia: DataUrodzenia{1, 2, 1990},
	}
	Jan := Osoba{
		"Jan",
		"Mielewczyk",
		170,
		DataUrodzenia{1, 10, 2000},
	}
	fmt.Println(Michał, Jan.DataUrodzenia)
	//	fmt.Println("dsd")
}
