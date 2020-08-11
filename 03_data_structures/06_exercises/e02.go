package main

import "fmt"

/* Exercise 02:
 * Create a slice with 10 elements
 * show every element with range loop
 * then show the slice type
 *
**/

func main() {
	slice := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}

	for key, value := range slice {
		fmt.Printf("%d:%d\n", key, value)
	}
	fmt.Printf("\nType: %T", slice)
}
