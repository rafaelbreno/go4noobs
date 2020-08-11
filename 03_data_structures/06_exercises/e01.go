package main

import (
	"fmt"
)

/*
 * Exercise 01:
 * Create an Array with 5 integers
 * and with a range, show for each its value and type
**/
func main() {
	arr := [5]int{3, 5, 7, 11, 13}

	for _, value := range arr {
		fmt.Printf("Value: %d - Type:%T\n", value, value)
	}

	fmt.Printf("\nArray type: %T", arr)
}
