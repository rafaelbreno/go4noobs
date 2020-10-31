package main

import (
	"fmt"
	"time"
)

func changeValue(i *int) {
	go func() {
		*i++
	}()
}

func printVar(i *int, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("Delay 2s", *i)
	}()
}

func main() {
	i := 10
	time2s, _ := time.ParseDuration("2s")

	printVar(&i, time2s)

	changeValue(&i)

	printVar(&i, time2s)

	c := time.Tick(500 * time.Millisecond)
	for _ = range c {
		fmt.Println(i)
	}
}
