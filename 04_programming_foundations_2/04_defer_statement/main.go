package main

import "fmt"

func defer01() {
	/* The 'defer fmt.Println("1st Line")'
	 * will be deferred
	 * That means, that line will be
	 * the last one printing a message
	 * Because of the 'defer' at the start of this line
	 */
	defer fmt.Println("1st Line")

	fmt.Println("2nd Line")
	fmt.Println("3rd Line")
	fmt.Println("4th Line")
}

func defer02() {
	/* What if we defer more than 1 line?
	 * Pretty simple, just need to understand this:
	 * "The first will be the last"
	 * So the "first defer", will be last executed of the "defers"
	 * and the "last defer", will be first executed of the "defers"
	 */
	defer fmt.Println("1st Line - 1st Defer")
	defer fmt.Println("2nd Line - 2nd Defer")
	defer fmt.Println("3rd Line - 3rd Defer")
	defer fmt.Println("4th Line - 4th Defer")
}

func main() {
	defer01()
	fmt.Println("------------------------")
	defer02()
}
