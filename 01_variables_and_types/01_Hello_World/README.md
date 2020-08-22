#### Hello World
- Like all tutorials we need to make a *Hello World* program
- In this example we can see the structure of a basic GO program
- ```go
      package main
      
      import (
        "fmt"
      )
      
      func main() {
        fmt.Println("Hello World,", "I'm Rafael Breno!")
      }
  ```
- When we run it with:
    - > go build main.go
    - output: "Hello World, I'm Rafael Breno!"