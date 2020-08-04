package main

import "fmt"

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
}
