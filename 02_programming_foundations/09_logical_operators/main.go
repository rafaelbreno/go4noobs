package main

import (
	"fmt"
	"strconv"
)

/*
* In this program I'll be developing
* a program that will build the same
* tables shown at README.md file
 */
func main() {
	var values [2]bool
	values[0] = false
	values[1] = true
	/*
	 * values = [false, true]
	 */
	fmt.Println("---And Operator---")
	for _, valueA := range values {
		for _, valueB := range values {
			fmt.Printf("%s && %s = %s\n",
				strconv.FormatBool(valueA),
				strconv.FormatBool(valueB),
				strconv.FormatBool(valueA && valueB))
		}
	}
	fmt.Println("---Or Operator---")
	for _, valueA := range values {
		for _, valueB := range values {
			fmt.Printf("%s || %s = %s\n",
				strconv.FormatBool(valueA),
				strconv.FormatBool(valueB),
				strconv.FormatBool(valueA || valueB))
		}
	}
	fmt.Println("---Not Operator---")
	for _, valueA := range values {
		fmt.Printf("!%s = %s\n",
			strconv.FormatBool(valueA),
			strconv.FormatBool(!valueA))
	}
}
