package main

import (
	"fmt"
)

/*
 * if (boolean condition) {
 * 		Do This
 * }
 * And then this
 */
/*
 * if (boolean condition) {
 * 		Do This
 * } else {
 *		If false, do this
 * }
 * And then this
 */
/*
 * if (boolean condition) {
 * 		Do This
 * } else if (boolean condition) {
 * 		If the one above if false
 * 		do this
 * } else {
 * 		if everything is false, do this
 * }
 * And then this
 */

var unknown interface{}

func main() {
	var numType string
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			numType = "even"
		} else {
			numType = "odd"
		}
		fmt.Println(i, " - ", numType)
	}

	switch numType {
	case "even":
		fmt.Println("The last number was an even number!")
		//fallthrough
	case "odd":
		fmt.Println("The last number was an odd number!")
		//fallthrough
	default:
		{
			fmt.Println("The values doesn't match any one of the possibilities")
		}
	}

	switch unknown.(type) {
	case int:
		fmt.Println("This is a integer!")
	case bool:
		fmt.Println("This is a boolean")
	case string:
		fmt.Println("This is a string")
	case float32, float64:
		fmt.Println("This is a float!")
	default:
		fmt.Println("Unknown type!")
	}
}
