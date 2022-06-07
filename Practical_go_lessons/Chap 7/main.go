package main

import "fmt"

func main() {
	//Hexadecimal representation
	// number := 270
	// fmt.Printf("%x\n", number)
	// fmt.Printf("%X\n", number)
	//adding 0x means hexadecimal
	// rep := 0x10e
	// fmt.Printf("%d\n", rep)
	//octal numeral system same but %o

	//bits
	b := make([]byte, 0)
	fmt.Println(b)
	b = append(b, 23)
	//b = append(b, 256) lare number
	fmt.Printf("%b\n", b)
	//Strings in go are IMMUTABLE
	// napis := "dsd"
	// napis = "ds"
	// fmt.Println(napis)

}
