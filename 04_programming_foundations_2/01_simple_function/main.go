package main

import "fmt"

/*
 * remember the shown function structure
 * in this case:
 *    - funcName = sum
 *    - paramName(paramType) = num1(int) and num2(int)
 *    - returnType = int
 */
func sum(num1 int, num2 int) int {
	/* This is the function scope
	 * the only things that can be accessed here is:
	 *     - num1
	 *     - num2
	 *     - global variables/constants
	 */

	// The return type is int, just the sum of the 2 numbers
	return num1 + num2
}

func main() {
	fmt.Println("5 + 9 =", sum(5, 9))
	fmt.Println("5 - 9 =", sum(5, -9))
	fmt.Println("10004638 + 128721671 =", sum(10004638, 128721671))
	//This won't work because it would be passing float values, not integer
	//fmt.Println("5.001 + 9.999 =", sum(5.001, 9.999))
}
