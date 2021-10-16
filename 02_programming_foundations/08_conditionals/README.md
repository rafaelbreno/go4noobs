#### Conditionals
- [Doc](https://golang.org/ref/spec#If_statements)
- Conditionals are decisions inside the code
```golang
package main

import "fmt"

func main() {
    var numType string
    for i := 1; i < 10; i++ {
        if i % 2 == 0 {
            numType = "even"
        } else {
            numType = "odd"
        }
        fmt.Println(i, " - ", numType)
    }
}
```
```golang
package main

import "fmt"

func main() {
    var numType string
    for i := 1; i < 10; i++ {
        if i%2 == 0 {
            numType = "even"
        } else {
            numType = "odd"
        }
    }

    switch numType {
        case "even": {
            fmt.Println("The last number was an even number!")
        } 
        case "odd": {
            fmt.Println("The last number was an odd number!")
        }
        default: {
            fmt.Println("The values doesn't match any one of the possibilities")
        }
    }
}
```
- The [switch](https://golang.org/ref/spec#Switch_statements) can be adapted to perform certain action depending on the variable type using _interfaces_
```golang
package main

import (
    "fmt"
)

var unknown interface{}

func main() {
      switch unknown.(type) {
      case int:
          fmt.Println("This is a integer!")
      case bool:
          fmt.Println("This is a boolean")
      case string:
          fmt.Println("This is a string")
      case float32, float64:
          fmt.Println("This is a float!")
      default:
          fmt.Println("Unknown type!")
      }
}
```
