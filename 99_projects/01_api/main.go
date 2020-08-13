package main

import (
	"fmt"
	"net/http"
)

const (
	url  = "http://localhost"
	port = "5000"
)

//
func concat(strings ...string) string {
	var concatStr string
	concatStr = ""
	for _, value := range strings {
		concatStr += value
	}
	return concatStr
}

func main() {
	/*
	 * This is just an example of a variadic function
	 * that will concatenate every string
	 * passed at the params
	 */
	basePort := concat(":", port)

	endPoint := "/"

	http.HandleFunc(endPoint, func(writer http.ResponseWriter, request *http.Request) {
		/* request.Method
		 * Will return a string
		 * of the type of HTTP Request
		 * sent to this endpoint
		 */
		method := request.Method

		/* Then just to test it
		 * It'll be printing the method
		 * so you can test it with Postman
		 */
		fmt.Println("Method ", method)
	})

	/*
	 * This will tell this app
	 * to start listening to a port
	 * so the app won't just end
	 * will stop here and listen to every
	 * connection at this port
	 */
	err := http.ListenAndServe(basePort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
