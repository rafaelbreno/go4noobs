package main

import (
	"fmt"
)

func main() {
	s := "Hello, World!"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T <-> %#U - %#x\n", v, v, v, v)
	}

	// Char per char
	for _, z := range s {
		fmt.Printf("%v - %T <-> %#U - %#x\n", z, z, z, z)
	}

	// Byte per Byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T <-> %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
