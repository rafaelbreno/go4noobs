package main

import (
	"fmt"
)

// type myType subType
type myType int

var b myType

func main() {
	a := 10
	fmt.Printf("%T", a)
	fmt.Printf("\n%T", b)
	/**
	* a = b this is wrong
	* because 'a' is type 'int'
	* and b is type 'myType'
	**/
}
