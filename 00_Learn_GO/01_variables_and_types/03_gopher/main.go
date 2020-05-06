package main

import (
	"fmt"
)

func main() {
	numInt := 20                           // %d
	numFloat := 10.2                       // %g
	textString := "Olá, meu nome é Rafael" // %s
	textChar := 'R'                        // %q
	fmt.Printf("Var: %d , Type: %T\n", numInt, numInt)
	fmt.Printf("Var: %g , Type: %T\n", numFloat, numFloat)
	fmt.Printf("Var: %s , Type: %T\n", textString, textString)
	fmt.Printf("Var: %q , Type: %T\n", textChar, textChar)
}
