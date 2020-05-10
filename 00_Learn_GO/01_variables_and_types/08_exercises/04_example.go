package main

import (
	"fmt"
)

type myType int

var a myType

func main() {
	fmt.Printf("Value: %v\n", a)
	fmt.Printf("Type: %T\n", a)
	a = 42
	fmt.Printf("Value: %v\n", a)
}
