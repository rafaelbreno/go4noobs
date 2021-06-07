package main

import (
	"fmt"
	"sync"
	"time"
)

type Hero struct {
	ID   uint
	Name string
}

// Func to add a hero to a channel
func sendHero(wg *sync.WaitGroup, h Hero, ch chan Hero) {
	// Increments the WaitGroup counter
	wg.Add(1)

	// Sleep 1 second to prove that the fmt.Print() will wait for the channel receive the next value
	time.Sleep(time.Second)

	// Adding hero to channel
	ch <- h
}

func main() {
	// Mocking a Hero list
	heroes := []Hero{
		{
			ID:   uint(1),
			Name: "Super Magnus",
		},
		{
			ID:   uint(2),
			Name: "McDevil(Not Evil)",
		},
		{
			ID:   uint(3),
			Name: "Big Chungus",
		},
		{
			ID:   uint(4),
			Name: "BASS Player",
		},
	}

	// Creating a WaitGroup
	var wg sync.WaitGroup

	// Creating a channel to receive Heroes
	heroChan := make(chan Hero, len(heroes))

	// Increments the WaitGroup counter
	// The GoRoutine below can send the first Hero before finishing the program
	wg.Add(1)

	// GoRoutine to iterate all Heroes
	go func() {
		for _, val := range heroes {
			// Sending heroes to Channel
			sendHero(&wg, val, heroChan)
		}

		wg.Done()
	}()

	// Iterate again the heroes
	go func() {
		for range heroes {
			// Will only print when the hero is add to channel
			fmt.Printf("%+v\n", <-heroChan)

			// Decrements the WaitGroup counter
			wg.Done()
		}
	}()

	wg.Wait()
}
