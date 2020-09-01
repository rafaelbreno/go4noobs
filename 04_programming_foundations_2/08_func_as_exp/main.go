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