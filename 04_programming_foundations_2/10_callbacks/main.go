package main

import (
	"fmt"
)

/* Function with a callback 'f'
 * as an Argument */
func separate(f func(num int) int, nums ...int) []int {
	var ns []int
	for _, value := range nums {
		// Here it'll be using f()
		i := f(value)
		if i != 0 {
			ns = append(ns, i)
		}
	}
	return ns
}

// This will return evens and zero
func even(n int) int {
	if n%2 == 0 {
		return n
	}
	return 0
}

// This will return odds and zero
func odd(n int) int {
	if n%2 != 0 {
		return n
	}
	return 0
}

func main() {
	var nums []int
	// Creating a slice from 1 to 21
	for i := 1; i <= 21; i++ {
		nums = append(nums, i)
	}
	// Showing all values
	fmt.Println("All values:", nums)

	/* Evens will receive the
	 * return value from
	 * separate(func even()m nums)
	 */
	evens := separate(even, nums...)
	fmt.Println("Evens values", evens)

	/* Odds will receive the
	 * return value from
	 * separate(func odd()m nums)
	 */
	odds := separate(odd, nums...)
	fmt.Println("Odd values", odds)
}
