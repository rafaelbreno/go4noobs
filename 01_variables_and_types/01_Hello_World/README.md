#### Hello World
- Like all tutorials we need to make a *Hello World* program
- In this example we can see the structure of a basic GO program
```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello World,", "I'm Rafael Breno!")
}
```
- We can run it with:
    - > go build main.go
        - Builds the file to a binary
    - output: "Hello World, I'm Rafael Breno!"
- Or we can just:
    - > go run main.go
- Now let's try to understand each part of this simple application:

##### `package main`
- Go applications are structure as packages, each folder, is a package
- Package are ways to define `namespaces` and `scopes`
  - You will see more about scopes later on
- The `namespace` is responsible to define the package name that will help us import later on
- Talking about `import`

##### `import`
- This is the way of referencing the Go's `standard library` and 3rd party packages
```
import (
  "fmt" // Here it's been imported the `fmt` library
)
```

##### `func main`
- `func` is the Go's function keyword, we'll be seing more about it later
- In Go, the application start on the `main()` function, so it'll be the headstart, the application's bootstrap

##### `fmt.Println`
- Here we're referencing the imported `fmt` package, and calling a function `Println`
