package main

import (
	"fmt"
)

func Sum(num1 int, num2 int) {
	num1 += num2
}

// Type will be a pointer of a int variable
func SumPtr(numPtr *int, num2 int) {
	*numPtr += num2
	// or
	// p = *numPtr
	// p += num2
}

func main() {
	num := 20

	// Initial value:  20
	fmt.Println("Initial value: ", num)

	Sum(num, 10)
	// Sum value(+10):  20
	fmt.Println("Sum value(+10): ", num)

	SumPtr(&num, 10)
	// SumPtr value(+10):  30
	fmt.Println("SumPtr value(+10): ", num)

	// Address value:  0xc000126010
	fmt.Println("Address value: ", &num)
}
