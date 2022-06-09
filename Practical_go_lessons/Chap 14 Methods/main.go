package main

import "fmt"

type Sea struct {
	Name     string
	Area     int
	Deep     int
	Capacity int
}

func (r Sea) NandC() (string, int) {
	return r.Name, r.Deep
}
func (r *Sea) CalculateCapacity() {
	r.Capacity = r.Area * r.Deep
}
func main() {
	Martwe := Sea{
		"Morze Martwe",
		10000,
		19999,
		0,
	}
	fmt.Println(Martwe.NandC())
	Martwe.CalculateCapacity()
	fmt.Println(Martwe.Capacity)
}

// !! Visible= Capitalized ,, Invisible==not capitalized
