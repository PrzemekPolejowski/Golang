package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.Open("przyklad.txt")
	scanner := bufio.NewScanner(data)
	caption := ""
	for scanner.Scan() {
		line := scanner.Text()
		single := strings.Split(line, " ")
		if single[0] == "DOPISZ" {
			caption += single[1]
		} else if single[0] == "ZMIEN" {
			var r rune
			r = 3
			out := []rune(caption)
			out[1] = r
			fmt.Println(out)
		}
	}
	fmt.Println(caption)
}
