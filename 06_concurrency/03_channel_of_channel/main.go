package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Method   chan string
	Endpoint string
}

func main() {
	var wg sync.WaitGroup

	methodChan := make(chan string, 4)

	reqs := []Request{
		{
			Method:   methodChan,
			Endpoint: "/login",
		},
		{
			Method:   methodChan,
			Endpoint: "/logout",
		},
		{
			Method:   methodChan,
			Endpoint: "/profile",
		},
		{
			Method:   methodChan,
			Endpoint: "/home",
		},
	}

	reqChan := make(chan Request, 4)

	wg.Add(1)

	defReq := func(req string, ch chan string, wg *sync.WaitGroup) {
		ch <- req
		wg.Add(1)
	}

	go func() {
		for _, req := range reqs {
			reqChan <- req
			if req.Endpoint == "/login" || req.Endpoint == "/logout" {
				defReq("POST", methodChan, &wg)
			} else {
				defReq("GET", methodChan, &wg)
			}
		}
		wg.Done()
	}()

	go func() {
		for range reqs {
			req := <-reqChan
			method := <-req.Method
			fmt.Printf("%s - %s\n", req.Endpoint, method)
			wg.Done()
		}
	}()

	wg.Wait()
}
