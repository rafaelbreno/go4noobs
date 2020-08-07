package main

import "fmt"

//                 0  1  2  3  4  5  6
var slice = []string{"a", "b", "c", "d", "e", "f", "g"}

func main() {
	fmt.Println("-----[:]-----")
	slice1 := slice[:]
	for key, value := range slice1 {
		fmt.Println(key, ":", value)
	}
	fmt.Println("\n-----[3:]-----")
	slice2 := slice[3:]
	for key, value := range slice2 {
		fmt.Println(key, ":", value)
	}
	fmt.Println("\n-----[:3]-----")
	slice3 := slice[:3]
	for key, value := range slice3 {
		fmt.Println(key, ":", value)
	}
	fmt.Println("\n-----[2:5]-----")
	slice4 := slice[2:5]
	for key, value := range slice4 {
		fmt.Println(key, ":", value)
	}

	fmt.Println("\n-----[:2] U [4:]-----")
	/*
	 * Again with math notations
	 * remember "U"?
	 * That means "union"
	 * In this case [:2] U [4:]
	 * Will be uniting the terms from 0 to 1, and from 4 until the last one
	 * To unite them both we'll be using append with the "..." operator AGAIN!
	 */
	union := append(slice[:2], slice[4:]...)
	fmt.Println(union)
}
