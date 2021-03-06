package main

import (
	"fmt"
	"time"
)

func changeValue(i *int) {
	// This func will simply add +1 to the variable
	go func() {
		*i++
	}()
}

func printVar(i *int, delay time.Duration) {
	go func() {
		/* The time.Sleep() function will pause
		 * this scope for "delay" time
		 * more at: https://golang.org/pkg/time/#Sleep
		 */
		time.Sleep(delay)
		// Then will print the value
		fmt.Printf("Delay %.0fs - %d\n", delay.Seconds(), *i)
		// showing the value modification after "delay" time
		*i++
	}()
}

func main() {
	i := 10
	time2s, _ := time.ParseDuration("2s")
	time5s, _ := time.ParseDuration("5s")

	// goroutine with 5s sleep
	printVar(&i, time5s)

	// changing value +1
	changeValue(&i)

	// goroutine with 2s sleep
	printVar(&i, time2s)

	/* This will run a scope each "x" time
	 * in this case 500 miliseconds (0,5 sec)
	 */
	c := time.Tick(500 * time.Millisecond)

	for range c {
		fmt.Println(i)
	}
}
