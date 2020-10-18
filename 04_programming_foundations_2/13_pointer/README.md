#### Pointers
- _"In computer science, a [Pointer](https://en.wikipedia.org/wiki/Pointer_(computer_programming)) is an object in many programming languages that stores a memory address."_
- For more info:
	- [GoByExample](https://gobyexample.com/pointers)
	- [Effective Go](https://golang.org/doc/effective_go.html#methods)
- Reading those two docs you will those characteres _"&"_ and _"*"_
	- __&__ - Is the address of a _variable_
	- __*__ - Will return a pointer of a _variable_
- 	```go
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
	```