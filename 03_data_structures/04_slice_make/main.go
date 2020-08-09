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
	fmt.Println("----Multidimensional----")
	/* This is a 3 dimensional array */
	var multid [][][]int
	/* That means every 'i', have and 'j' of 'k'
	 * If you're already familiarized with math you know about R³(or R3)
	 * This means there will be a "point" with 3 coordinates (i, j, k)
	 * and this point will have a value
	 */

	/* The example below we'll be building
	 * a slice with 3 "cardinal points"
	 * (i, j, k) and the value will be the sum of 3 of them
	 * f(i, j, k) = i + j + k
	**/

	/* Here it will create an 'i' point
	 * and assign a slice [][]int{} value to it
	 */
	for i := 0; i < 3; i++ {
		multid = append(multid, [][]int{})

		/* Here it will create a 'j' point
		 * and assign a slice []int{} value to it
		 */
		for j := 0; j < 3; j++ {
			multid[i] = append(multid[i], []int{})

			/* Here will create the 'k' point
			 * and assign a integer value to it
			**/
			for k := 0; k < 3; k++ {
				multid[i][j] = append(multid[i][j], i+j+k)
			}
		}
	}

	// Accessing i slice([][]int{})
	for i_pos, i := range multid {
		// Accessing j slice([]int{})
		for j_pos, j := range i {
			// Accessing k value(int)
			for k_pos, k := range j {
				// k_pos is the position of value k
				fmt.Printf("(%d, %d, %d) = %d\n", i_pos, j_pos, k_pos, k)
			}
		}
	}
}
