package main

import (
	"fmt"
)

//This is a "global" variable
var global = 10

func main() {
	local := 20
	fmt.Print("Global Variable:", global)
	fmt.Print("\nLocal Variable:", local)
	Scope(local)
}

func Scope(x int) {
	//This is legal
	fmt.Print("\n\nGlobal Variable:", global)
	//This is legal
	fmt.Print("\nParameter Variable:", x)
	//This is legal, will show 'undefined' error in terminal
	// fmt.Print("Main Variable:", local)
}
