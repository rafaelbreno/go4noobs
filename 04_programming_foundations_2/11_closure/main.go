package main

import "fmt"

func main() {

	/* This will instantiate a new
	 * function scope */
	a := clos(10)

	fmt.Println(a()) // 11
	fmt.Println(a()) // 12
	fmt.Println(a()) // 13

	/* This will instantiate a new
	 * function scope */
	b := clos(245)

	fmt.Println(b()) // 246
	fmt.Println(b()) // 247
	fmt.Println(b()) // 248
}

/* This function will return a function and it's scope
 * inside of it's scope will have a variable "x"
 *		that receives a value of "i"
 * and then it will return a function that everytime it's called
 * will increment by one the variable "x"
 */
func clos(i int) func() int {
	x := i
	return func() int {
		x++
		return x
	}
}
