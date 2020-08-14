package main

import "fmt"

// The same example, but now a variadic version
func sum(nums ...int) int {
	/* It will form a slice of int
	 */
	fmt.Printf("Type: %T\nItems: %d\n", nums, nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	/* Sending the values
	 * If I wanted I could add even more number
	 * a finite amount like, 10, 40, 100, 1000 values
	 * and just that little function there
	 * would handle everything
	 */
	result := sum(1, 2, 3, 4, 5, 6)

	fmt.Println("\n\n1 + 2 + 3 + 4 + 5 + 6 = ", result)
}
