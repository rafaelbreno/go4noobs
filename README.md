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
        - [Twitter](https://twitter.com/ellenkorbes) || [GOLang Playlist<small>(in pt-br)</small>](https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)
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
     func main() {
     	//Here is  simple function
     	bytes, err := fmt.Println("Hello World,", "I'm Rafael Breno!")
     	fmt.Println("The string has", bytes, "bytes. Error:", err)
     }
  ```
- How do I know what's returning, and in what order? Following the doc:
    - "*It returns the number of bytes written and any write error encountered.*"
#### 03_Gopher
- Gopher is a short operator Its symbol is *":="* 
- Different of *"="*, that is a *assignment operator*
- You can see more of it in these links:
    - [Official Doc - Short Variable Declarations](https://golang.org/ref/spec#Short_variable_declarations)
    - [Unofficial Doc - Stackoverflow Question](https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go/45654233#45654233)
- At the Stackoverflown question, I've founded these *"rules"* about the **":="**
    1. You can't use *":="*  out of *"funcs"* . It's because, out of any func, a statement should start with a keyword.
        - ```go
          // no keywords below, illegal.
          illegal := 42
          
          // `var` keyword makes this statement legal.
          var legal = 42
          
          func foo() {
            alsoLegal := 42
            // reason: it's in a func scope.
          }
          ```
    2. You can't use them twice (in the same scope):
        - ```go
          legal := 42
          legal := 42 // <-- error
          ```
    3. You can use them for multi-variable declarations and assignments:
        - ```go
          foo, bar   := 42, 314
          jazz, bazz := 22, 7
          ```
    4. You can use them twice in **"multi-variable"** declarations, if one of the variables is new:
        - ```go
          foo, bar  := someFunc()
          foo, jazz := someFunc()  // <-- jazz is new
          baz, foo  := someFunc()  // <-- baz is new
          ```
    5. You can use the short declaration to declare a variable in a newer scope even that variable is already declared with the same name before:
        - ```go
          var foo int = 34
          
          func some() {
            // because foo here is scoped to some func
            foo := 42  // <-- legal
            foo = 314  // <-- legal
          }
          ```
    6. You can declare the same name in short statement blocks like: **if**, **for**, **switch**:
        - ```go
          foo := 42
          if foo := someFunc(); foo == 314 {
            // foo is scoped to 314 here
            // ...
          }
          // foo is still 42 here
          ```
- With it said, we can now use the function [Printf]() and the [Printing "verbs"](https://golang.org/pkg/fmt/#hdr-Printing)
- To help understand that *":="* guess what type of value it's assigned to it
- ```go
  func main(){
  	numInt := 20
  	numFloat := 10.2
  	textString := "Olá, meu nome é Rafael"
  	textChar := 'R'
  	fmt.Printf("Var: %b , Type: %T\n", numInt, numInt)
  	fmt.Printf("Var: %b , Type: %T\n", numFloat, numFloat)
  	fmt.Printf("Var: %b , Type: %T\n", textString, textString)
  	fmt.Printf("Var: %b , Type: %T\n", textChar, textChar)
  }
  ```
