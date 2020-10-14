package main

import "fmt"

func main() {
	fmt.Println(fac(5))
}

// Factorial
func fac(n int) int {
	// The function "breaks" when n == 1
	if n == 1 {
		return 1
	}
	return fac(n-1) * n
}
