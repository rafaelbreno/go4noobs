package main

import (
	"fmt"
	"sync"
	"time"
)

// This will allow us to great those waiting groups
var waitGroup sync.WaitGroup

func main() {
	/* Here we can add the amount of goroutines
	 * that we'll be working on
	 * this case I want to great 2 routines
	**/
	waitGroup.Add(2)

	go awaitFunc("5s")
	go awaitFunc("2s")
	/* This will say to our program to not stop
	 * And wait for all goroutines to end their processes
	 * before ending/exiting the program
	**/
	waitGroup.Wait()
}

func awaitFunc(timeout string) {
	/* Like we've seen before
	 * this will stop the function for some amount of time
	 * just to simulate a long process
	**/
	delay, _ := time.ParseDuration(timeout)

	time.Sleep(delay)
	fmt.Printf("Have been passed %s!\n", timeout)

	/* This will send a message to our WaitGroup
	 * "Hey this process have finished"
	 * and depending on how many process we've added
	 * to the wait group, the func main() will end or not
	**/
	waitGroup.Done()
}
