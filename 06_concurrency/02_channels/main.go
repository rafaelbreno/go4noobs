package main

import (
	"fmt"
	"time"
)

func main() {
	// Making a channel
	c := make(chan int)

	time2s, _ := time.ParseDuration("4s")

	// Goroutine
	go func() {
		time.Sleep(time2s)
		/* After running some tasks
		 * it's returned value 1
		 */
		c <- 1
	}()

	tick := time.Tick(500 * time.Millisecond)

	for _ = range tick {
		// Goroutine
		go func() {
			// Waiting for channel end it's tasks
			if <-c == 1 {
				fmt.Printf("Passou %.0fs\n", time2s.Seconds())
			}
		}()

		fmt.Println("a")
	}
}
