/*
 * Write expressions using the following operators and assign its values to variables
	- ==
	- !=
	- >=
	- >
	- <=
	- <
*/
package main

import "fmt"

func main() {
	var arr [6]bool

	arr[0] = 10 == 20
	arr[1] = 10 != 20
	arr[2] = 10 >= 20
	arr[3] = 10 > 20
	arr[4] = 10 <= 20
	arr[5] = 10 < 20

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, " - ", arr[i])
	}

}
