#### 01 Variables and Types
- This section is an introduction to GO
1. Hello World
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
2. Println()
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
3. Gopher
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
4. Scope
    - Scope is an area with limits defined by ``{}``
        - For Example``{ ...in scope... } ...out of scope`` 
    - A good example of it is a "Global Variable", its **scope** is the entire code, can be used in every *func* of a file
    - At the code below we can see that the variable **global** can be used 
    - ```go
          var global = 10
          func main() {
            local := 20
            fmt.Print("Global Variable:", global)
            fmt.Print("\nLocal Variable:", local)
            Scope(local)
          }     
          func Scope(x int){
            //This is legal
            fmt.Print("\n\nGlobal Variable:", global)
          
            //This is legal
            fmt.Print("\nParameter Variable:", x)
          
            //This is illegal, will show 'undefined' error in terminal
            // fmt.Print("Main Variable:", local)
          }
      ```
5. Types
    - The title is self-explanatory, type is the *"format"* of a designed variable
    - Assigning a type for a variable it means that variable can only accept it until the end of the program
    - These types can be:
        - Primitive
            - | Title   | Type    | %  | Example   |
              |---------|---------|----|-----------|
              | Integer | int     | %b | 10        |
              | Float   | float32 | %b | 10.2      |
              | Boolean | bool    | %t | true      |
              | String  | string  | %t | "Example" |
        - [Derived](https://www.tutorialspoint.com/go/go_data_types.htm)
            - Pointer
            - Array
            - Structure
            - Union
            - Function
            - Slice
            - Interface
            - Map
            - Channel
    - For now I'm going to stick with only *Primitive* types
    - ```go
          var a = 10
          var b float32
          /*
          Can't assign a value after declaring a variable out of a func
          b = 20
          */
          
          func main()  {
            b = 20.147
            var c bool = true
            var d string = "Rafael"
            fmt.Printf("Var: %v , Type: %T\n", a, a)
            fmt.Printf("Var: %v , Type: %T\n", b, b)
            fmt.Printf("Var: %v , Type: %T\n", c, c)
            fmt.Printf("Var: %v , Type: %T", d, d)
          }
      ```
