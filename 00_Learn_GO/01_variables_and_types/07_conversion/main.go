package main

import (
	"fmt"
)

//Creating a type based on 'int'
type myType int

var b myType = 20

func main() {
	a := 10
	fmt.Println("Passing value from b to a:")
	fmt.Printf("b myType -> Type %T , Value %v\n", b, b)
	a = int(b)
	fmt.Printf("a int -> Type %T , Value %v\n", a, a)
	fmt.Println("Passing value from a to b:")
	b = myType(a)
	fmt.Printf("b myType -> Type %T , Value %v\n", b, b)
	fmt.Printf("a int -> Type %T , Value %v\n", a, a)
}
