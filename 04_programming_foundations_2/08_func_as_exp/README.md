#### Function as Expression/Variable
- [_sorry didn't found an official Doc_]
- 	```go
		package main

		import (
			"fmt"
		)

		func main() {
			/* You can define a function
			 * as a variable
			 * and call it inside the scope
			 * that it's declared
			*/
			g := func(nums ...int) int {
				sum := 0
				for _, value := range nums {
					sum += value
				}
				return sum
			} 
			/* Here we can sum the values
			 * with the g()
			 * its the same as:
			func sum(nums ...int) int {
				sum := 0
				for _, value := range nums {
					sum += value
				}
				return sum
			} 
			* But func as expression/variable
			* are only accessible inside the scope
			* that it was declared 
			*/
			fmt.Println("Sum: ", g(1, 2, 3, 4, 5))
		}
	```
- Declaring a _function_ as variable will limit it
	- It will only be available inside the [_Scope_](https://github.com/rafaelbreno/go4noobs/tree/master/01_variables_and_types/04_Scope) that it's declared