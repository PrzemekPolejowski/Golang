package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	year := time.Now().Year()
	fmt.Println(year)
	day := time.Now().Day()
	fmt.Println(day)
}
