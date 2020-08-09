package main

import "fmt"

func main() {
	/* Just explaining the map structure
	 * map[key]value
	 * the 'key' is what is searched
	 * 'value' is what is returned
	**/
	intString := make(map[int]string, 10)

	// 02_programming_foundations/02_string
	hello := []byte("hello")

	for key, value := range hello {
		intString[key] = string(value)
	}
	/* The output will follow the structure
	 * key:value , almost like JSON
	 * [0:h 1:e 2:l 3:l 4:0]
	**/
	fmt.Println(intString)
}
