package main

import "fmt"

func multipleReturn(num int, words ...string) (int, string) {
	/* The first difference is at returnType
	 * Before we've seen only 'string', or 'int'
	 * In this case, we can add n-types inside the ()
	 */
	/* This function will be:
	 * Factorial
	 * String concatenate
	 */
	fac := 1
	for i := 1; i <= num; i++ {
		fac *= i
	}
	phrase := ""
	for _, word := range words {
		phrase += word + " "
	}

	/* Here is another difference
	 * like the returnType, we'll be return
	 * the values separated by a comma ","
	 * Remember:
	 * FOLLOWING THE SAME ORDER OF returnType
	 */
	return fac, phrase
}

func main() {
	//  (int, string) remember the order?
	factorial, sentence := multipleReturn(8, "Hello!", "This", "is", "a", "multiple", "return", "function")

	fmt.Println(sentence)
	fmt.Println("Factorial:", factorial)
}
