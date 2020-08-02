package main

import "fmt"

const (
	a = iota // iota == 0
	b = iota // iota == 1
	c = iota // iota == 2
	d = iota // iota == 3
)

func main() {
	fmt.Println(a, b, c, d)
	str := ""
	_, _ = fmt.Scanln(str)
	fmt.Println("- ", str, " - ")
}
