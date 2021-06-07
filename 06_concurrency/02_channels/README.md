### Channels
- [Doc](https://golang.org/doc/effective_go.html#channels)
- Channels are create with _make(chan type, size)_
    - The same as we've seen with [maps](https://github.com/rafaelbreno/go4noobs/tree/master/03_data_structures/05_maps)
- So to create one we could simply: ```channel := make(chan int)```
-   ```go
        
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
    ```

- That was a simple example, but there's an important topic about it
    - The channel will pause the execution of a program, until it have received the value of it
    - This is known as _channel syncronization_
- You can see it in the example below:
```golang
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
```
