package main

import "fmt"

/* Don't need to be scared
 * TODO: Add chapter url
 * We'll be studying function at Chapter
 * func funcName(params) returnType {
 * 		func scope
 * }
**/

/* Name: true
 * Params: nothing
 * Return Type: String
 * Return Value: "true"
 */
func returnTrue() string {
	return "True"
}

/* Name: true
 * Params: nothing
 * Return Type: String
 * Return Value: "true"
 */
func returnFalse() string {
	return "False"
}

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

	/* Here's a little more "advanced" example
	 * we'll be using functions
	 * Don't need to be scared, we'll be entering in that chapter
	 * Here:
	 * FOR NOW You only need to know that a function, is a structure
	 * that returns something
	**/
	funcs := map[int]func() string{
		//key: function
		0: returnFalse,
		1: returnTrue,
	}

	// key, value := range map
	for _, f := range funcs {
		/* r will be the return valued from a function
		 * Again, don't need to get frighten about function
		 * we'll be studying it more deeply again at chapter
		 * TODO: add function chapter link
		**/
		r := f()
		fmt.Println(r)
	}
}
