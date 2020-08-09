package main

import "fmt"

func main() {

	fmt.Println("----Len < Cap")
	slice := make([]int, 3, 5)
	for key, value := range slice {
		fmt.Println(key, "-", value)
	}
	fmt.Println("----Overflowing Cap----")
	for i := 1; i <= 5; i++ {
		slice = append(slice, i)
	}
	for key, value := range slice {
		fmt.Println(key, "-", value)
	}
	/*
	 * When appending, you can overflown the capacity
	 * because slices are flexible and expandable
	 */
}
