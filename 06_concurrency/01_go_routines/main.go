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
		fmt.Printf("Delay %.0fs - %d\n", delay.Seconds(), *i)
		*i++
	}()
}

func main() {
	i := 10
	time2s, _ := time.ParseDuration("2s")
	time5s, _ := time.ParseDuration("5s")

	printVar(&i, time5s)

	changeValue(&i)

	printVar(&i, time2s)

	c := time.Tick(500 * time.Millisecond)
	for _ = range c {
		fmt.Println(i)
	}
}
