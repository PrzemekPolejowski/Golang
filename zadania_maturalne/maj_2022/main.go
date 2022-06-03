package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var liczbyPierwsze [9592]int

func pierwsze() {
	var counter = 0
	for i := 2; i < 100000; i++ {
		exit := false
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				exit = true
				break
			}
		}
		if !exit {
			liczbyPierwsze[counter] = i
			counter++
		}
	}

}
func zad1() {
	//Zad 1
	data, _ := os.Open("liczby.txt")
	scanner := bufio.NewScanner(data)
	counter := 0
	help := false
	for scanner.Scan() {
		line := scanner.Text()
		first, last := string(line[0]), string(line[len(line)-1])
		if first == last {
			if !help {
				fmt.Print(line, " ")
				help = true
			}
			counter++
		}
	}
	fmt.Println(counter)
}
func zad2() {
	pierwsze()
	data, _ := os.Open("liczby.txt")
	scanner := bufio.NewScanner(data)
	counter_max := 0
	value_max := 0
	second_counter_max := 0
	second_value_max := 0
	for scanner.Scan() {
		line := scanner.Text()
		single, _ := strconv.Atoi(line)
		counter := 0
		second_counter := 0
		kopia := single
		var pomoc [100000]int

		for single != 1 {
			for i := 0; i < 9591; i++ {
				//fmt.Println(liczbyPierwsze[i])
				if single%liczbyPierwsze[i] == 0 {
					if pomoc[liczbyPierwsze[i]] == 0 {
						pomoc[liczbyPierwsze[i]] = 1
						second_counter++
						if second_counter > second_counter_max {
							second_counter_max = second_counter
							second_value_max = kopia
						}
					}
					counter++
					if counter > counter_max {
						counter_max = counter
						value_max = kopia
					}
					single /= liczbyPierwsze[i]
					break
				}
			}

		}
	}
	fmt.Println(value_max, counter_max, second_value_max, second_counter_max)
}
func zad3() {
	data, _ := os.Open("liczby.txt")
	scanner := bufio.NewScanner(data)
	var tab [200]int
	counter := 0
	zliczacz_trojek, zliczacz_piatek := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		single, _ := strconv.Atoi(line)
		tab[counter] = single
		counter++
	}
	for _, i := range tab {
		for _, j := range tab {
			if j != i && j%i == 0 {
				for _, z := range tab {
					if z != j && z%j == 0 {
						for _, c := range tab {
							if c != z && c%z == 0 {
								for _, n := range tab {
									if n != c && n%c == 0 {
										//fmt.Println(i, j, z, c, n)
										zliczacz_piatek++
									}
								}
							}
						}
						//fmt.Println(i, j, z)
						zliczacz_trojek++
					}
				}
			}
		}
	}
	println(zliczacz_trojek, zliczacz_piatek)
}
func main() {
	fmt.Print("Zad 1: ")
	zad1()
	fmt.Print("Zad 2: ")
	zad2()
	fmt.Print("Zad 3: ")
	zad3()
}
