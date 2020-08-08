package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)

	for key, value := range slice {
		fmt.Println(key, "-", value)
	}
}
