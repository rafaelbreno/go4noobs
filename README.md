# DevOps
## Summary
- 👨‍💻 👩‍💻 In development 🚧 🛠️
## Introduction
- The purpose of this repo is to help me track my evolution in:
    - DevOps
    - Go Lang
- To help me find a "north" I'm going to use as a *Guide* the site [RoadMap](https://roadmap.sh/devops)
- During this long process I'm will be documenting all topics, details, tips, **everything**
- So this maybe(focus on **MAYBE**) can be a "guide" to you too
## Go Lang
- Like I've said above, I'm going to use GoLang to accompany me go through this learning process
- If you have never heard of *Go* before, this links may help you understand:
    - [Official Documentation](https://golang.org/doc/)
    - [Is Go OOP?](https://golang.org/doc/faq#Is_Go_an_object-oriented_language) <small><small>no and yes ???</small></small>
    - [Principles](https://golang.org/doc/faq#principles)
    - A special thanks to *Ellen Körbes*, for bringing an amazing GOLang content to Brazil in PT-BR
        - [Twitter](https://twitter.com/ellenkorbes) || [GOLang Playlist<small>In pt-br</small>](https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)
- I'm using for IDE [GOLand](https://www.jetbrains.com/go/) from [JetBrains](https://www.jetbrains.com/)
- And a pretty good alternative is [Go Playground](https://play.golang.org/), an online IDE
- With everything setup, we're ready to start programming, Let's *GO*!!!
---------
# [Roadmap](https://roadmap.sh/devops)
## 00 - Learn a Programming Language - **GO**
### 01 - Variables and types
- This section is an introduction to GO
#### 01 - Hello World
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
#### 02 - Println()
- I think it's a good habit read the [Official Documentation](https://golang.org/doc/) of any lang/framework/tool that you're using
- This case is a good example of that
- Here is the [Println() Func Doc](https://golang.org/pkg/fmt/#Println), reading its possible to do this:
- ```go
     package main
     
     import (
     	"fmt"
     )
     
     func main() {
     	//Here is  simple function
     	bytes, err := fmt.Println("Hello World,", "I'm Rafael Breno!")
     	fmt.Println("The string has", bytes, "bytes. Error:", err)
     }
  ```
- How do I know what's returning, and in what order? Following the doc:
    - "*It returns the number of bytes written and any write error encountered.*"
