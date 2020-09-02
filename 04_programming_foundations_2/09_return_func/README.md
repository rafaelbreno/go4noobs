#### Returning a Function
- It's almost like Funcs as expression
- but _instead_ it's a function that returns other functions
- The example below will be a simple _math()_ function
- It'll have 2 options
	- Sum and Multiplication
	- Both will be [variadic functions](https://github.com/rafaelbreno/go4noobs/tree/master/04_programming_foundations_2/02_variadic_functions)
- The structure will be pretty simple
- 	```
		func funcName(funcArguments) func (returnedFuncArguments) TypeOfReturnedFunc {
			funcScope
			return returnedFunc (returndcFuncArguments) TypeOfReturnedFunc {
				returnedFuncScope
			}
		}
	```
- Just remembering Go is a _typed_ programming language
- The shown structure will look kinda strange
	- funcName - is the name of the functions that will return other functions
	- funcArguments - is the arguments of the function that will return other functions
	- returnedFuncArguments - are the arguments of the returned functions
	- TypeOfReturnedFunc - is the value type that the returnedFunc will return
- Explaining like this is kinda(totally) confusing, but the example below will help you understand 
- 	```go
		package main
		import (
			"fmt"
		)
		
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

			fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
			fmt.Println("Multiplication:", mult(1, 2, 3, 4, 5))
		}
	```