/*
 * Using a loop
 * print a value in the following formats
 *		- Decimal 		%d
 *		- Hexadecimal 	%#x
 *		- Unicode 		%#U
 *		- Tab 			\t
 *		- New Line 		\n
 * Values from 33 to 122
 *
 */
package main

import "fmt"

func main() {
	var value int
	value = 33

	for value <= 122 {
		fmt.Printf("%d\t%#x\t%#U\n", value, value, value)
		value++
	}
}
