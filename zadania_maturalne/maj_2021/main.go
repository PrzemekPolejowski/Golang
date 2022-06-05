package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	data, _ := os.Open("przyklad.txt")
	scanner := bufio.NewScanner(data)
	var caption []byte
	for scanner.Scan() {
		line := scanner.Text()
		single := strings.Split(line, " ")
		if single[0] == "DOPISZ" {
			caption = append(caption, byte(single[1]))
		} else if single[0] == "ZMIEN" {
			//Problem
			// var help byte
			// help = single[1]
			// caption[len(caption)-1] = single[1]
			//Problem
		} else if single[0] == "USUN" {

		} else {

		}
	}
	//fmt.Println(caption)
}
