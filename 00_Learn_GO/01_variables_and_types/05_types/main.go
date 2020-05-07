package main

import (
	"fmt"
)

var a = 10
var b float32

/*
Can't assign a value after declaring a variable out of a func
b = 20
*/

func main() {
	b = 20.147
	var c bool = true
	var d string = "Rafael"
	fmt.Printf("Var: %v , Type: %T\n", a, a)
	fmt.Printf("Var: %v , Type: %T\n", b, b)
	fmt.Printf("Var: %v , Type: %T\n", c, c)
	fmt.Printf("Var: %v , Type: %T", d, d)
}
