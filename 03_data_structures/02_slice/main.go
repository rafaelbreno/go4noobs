package main

import "fmt"

var slice1 []int
var slice2 []int

var sliceBase []int
var sliceBase2 []int
var sliceBase3 []int
var sliceAdd = []int{1, 2, 3, 4, 5, 6}

func main() {
	/* append will a insert a value at the end of the slice
	 * Imagine that a slice is a queue of items
	 * appending something will put it at the end of the queue
	 */
	fmt.Println("----prepend----")
	for i := 0; i < 5; i++ {
		// This for will append values from 0 to 4
		slice1 = append(slice1, i)
	}
	for _, value := range slice1 {
		// Here will show the items at the order they were appended
		fmt.Println(value)
	}

	/*
	 * What if I wanted to prepend the values, for some reason
	 * "prepend" a value is putting the item at the beginning of the queue
	 */
	fmt.Println("----prepend----")
	for i := 0; i < 16; i++ {
		if i%2 == 0 {
			// If is a even number
			slice2 = append(slice2, i)
		} else {
			// If is a odd number
			slice2 = append([]int{i}, slice2...)
		}
	}
	fmt.Println(slice2)
	/*
	 * how that works?
	 * append is a function that receive 2 params
	 * append(slice, element) = [slice, element]
	 * so, to create a prepend just invert the order
	 * append(element, slice), but in this case, element isn't a slice type, and slice isn't a element
	 * to convert them we do this
	 * append([]int{element}, slice...)
	 * "slice..." will transform slice into element
	 * "[]int{element}" will transform element in a slice of int with unique item element
	 * kinda confusing but that makes sense
	 */

	fmt.Println("\n-----Base1-----")
	for _, value := range sliceAdd {
		sliceBase = append(sliceBase, value)
	}
	fmt.Println(sliceBase)

	fmt.Println("\n-----Base2-----")
	/*
		Look here
		append([]interface{}, 1, 2, 3, 4, ..., n - 1, n)
	*/
	sliceBase2 = append(sliceBase2, 1, 2, 3, 4, 5, 6)
	fmt.Println(sliceBase2)

	fmt.Println("\n-----Base3-----")
	sliceBase3 = append(sliceBase3, sliceAdd...)
	fmt.Println(sliceBase3)
}
