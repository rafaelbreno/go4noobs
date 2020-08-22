package main

import (
	"fmt"
)

func main() {
	/* Based on some variadic functions
	 * now we'll be concatenating some strings
	 */
	// First of all I'll be creating a slice of strings
	strs := []string{"Foo", "Bar", "Example", "Hello"}

	/* There's a slightly difference between:
	 * 		- "Normal" Func
	 * 		- Literal func
	 * here We can see the structure of an Literal func
	 * func(declared Params) (returnType) {
			... func scope ...
	 *	}(received Params)
	 * The first difference: There's no name, we don't declare it
	 * Second Difference: If the function need some params, just "input" it at the end of the func
	*/
	// variable       := func(params) (returnType) {
	concatenatedStr := func(words ...string) string {
		// function Scope
		str := ""
		for _, value := range words {
			str += value + " "
		}
		return str
	}(strs...)
	// ...end of func scope...}(receivedParams)

	/* There's possible create a
	 * function without a return type
	 * The example below We'll be sending the variables
	 * and then just outputting the received values
	 */
	func(sliceString []string, concatStr string) {
		fmt.Println("Slice of Strings:", sliceString)
		fmt.Println("\nConcatenated Strings:", concatStr)
	}(strs, concatenatedStr)
}
