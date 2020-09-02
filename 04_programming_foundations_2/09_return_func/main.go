package main

import (
	"fmt"
)

/* Return Function
 * It's almost like Funs as expression
 * but instead it's a function that return
 * functions as expression
*/

/* The example will be
 * based on a simple function called "math"
 * This function "math" will return other variadic functions
 * "sum" and "multiplication"
 * */
func math (opt int) func(...int) int {
	// Return default function
	var f = func (nums ...int) int {
		return 2
	}
	switch opt {
		// Sum
		case 1: 
			f = func (nums ...int) int {
				sum := 0
				for _, value := range nums {
					sum += value
				}
				return sum
			}
		// Multiplication 
		case 2:
			f = func (nums ...int) int {
				mult := 1
				for _, value := range nums {
					mult *= value
				}
				return mult
			}
	}
	// Returns default functions
	return f
}

func main() {
	// opt = 1
	sum := math(1)
	// opt = 2
	mult := math(2)

	// Can be done in both ways
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5), " - ", math(1)(1, 2, 3, 4, 5))
	fmt.Println("Multiplication:", mult(1, 2, 3, 4, 5), " - ", math(2)(1, 2, 3, 4, 5))
}