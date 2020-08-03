package main

import "fmt"

func main() {
	fmt.Printf(varTest)   // var_test doesn't exist
	varTest := "Teste 01" // var_test being declared
	fmt.Printf(varTest)   // var_test do exist
}
