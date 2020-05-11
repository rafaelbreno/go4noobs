package main

import (
	"fmt"
)

type myType int

var x myType
var y int

func main() {
	fmt.Printf("Value X: %v\n", x)
	fmt.Printf("Type X: %T\n", x)
	x = 42
	fmt.Printf("Value X: %v\n", x)
	y = int(x)
	fmt.Printf("Value Y: %v\n", y)
	fmt.Printf("Type Y: %T\n", y)
}
