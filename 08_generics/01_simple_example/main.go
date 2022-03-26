package main

import "fmt"

func Ptr[T any](value T) *T {
  return &value
}

func main() {
	v := "some value"
	fmt.Println(Ptr(v))
}
