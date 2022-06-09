package main

import (
	//"database/sql"
	"fmt"
)

//Init function run before main, with init function we can initialize a package
// func init() {
// 	fmt.Println("first")

// Key takeaways:
// Initialization of packages takes place in init functions.

// init functions are not mandatory.

// A package can declare several init functions.

// init functions are called after variable initialization.

// init functions are run sequentially.
// }
func init() {
	//sql.Register("mysql")
}

func main() {
	fmt.Println("sds")
}
