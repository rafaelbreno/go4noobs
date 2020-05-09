package main

import (
	"fmt"
)

func main() {
	//Here is  simple function
	bytes, err := fmt.Println("Hello World,", "I'm Rafael Breno!")
	fmt.Println("The string has", bytes, "bytes. Error:", err)
}
