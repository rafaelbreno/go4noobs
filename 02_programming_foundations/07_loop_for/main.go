package main

import "fmt"

func main() {
	// Types of for
	var i int
	var n []int
	fmt.Print("\n------|Do While i < 3|------\n")
	for i < 3 {
		fmt.Println("Loop 1: ", i)
		i = i + 1
	}
	fmt.Print("\n------|Common For|------\n")
	for i = 0; i < 5; i++ {
		fmt.Println("Loop 2: ", i)
	}
	i = 0
	// Do While i != 3
	fmt.Print("\n------|Do While i != 3|------\n")
	for {
		if i == 3 {
			break
		}
		fmt.Println("Loop 3: ", i)
		i++
	}
	fmt.Print("\n------|Only shows n + 1|------\n")
	i = 0
	for i < 6 {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println("Loop 4: ", i)
	}

	for i = 1; i < 6; i++ {
		n = append(n, i)
	}

	fmt.Print("\n------|Foreach For|------\n")
	for i = 0; i < len(n); i++ {
		fmt.Println("Foreach 1: ", n[i])
	}
	fmt.Print("\n------|Foreach Range|------\n")
	/*
	 * key, value := range arr
	 */
	for key, value := range n {
		fmt.Printf("%x: %x\n", key, value)
	}
}
