/*
 * Problem 01:
 *	Write a program that shows a number in:
	- decimal, binary and hexadecimal
*/
package main

import "fmt"

func main() {
	var num int
	// Decimal - Base 10
	fmt.Printf("%d", num)
	// Binary - Base 2
	fmt.Printf("%b", num)
	// Hexadecimal - Base 16
	fmt.Printf("%#x", num)
}
