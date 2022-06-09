package main

import (
	"fmt"
	"sort"
)

func main() {
	tab := []int{168, 373, 273, 390, 230, 444, 163, 10, 481, 8, 298, 304} //Sortowanie intow
	fmt.Println(tab)
	sort.Ints(tab)
	fmt.Println(tab)
	tabString := []string{"zoo", "monday", "jutro", "slice"} //Sortowabue stringow
	fmt.Println(tabString)
	sort.Strings(tabString)
	fmt.Println(tabString)
	//Sotrowanie map
	mapa := make(map[string]int)
	mapa["zoo"] = 1
	mapa["Jan"] = 2
	mapa["Alfabet"] = 3
	keys := []string{}
	for k := range mapa {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, mapa[v])
	}
}
