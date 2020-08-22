#### Defer Statement
- [Doc](https://golang.org/ref/spec#Defer_statements)
- _"A ['defer'](https://golang.org/doc/effective_go.html#defer) statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body..."_
- This means that a when you append the statement _defer_ at a line, that line will only be executed/ran at the end of that scope
-   ```go
        package main
        
        import "fmt"
        
        func main()  {
            /* The 'defer fmt.Println("1st Line")'
             * will be deferred
             * That means, that line will be
             * the last one printing a message
             * Because of the 'defer' at the start of this line
             */
            defer fmt.Println("1st Line")
        
            fmt.Println("2nd Line")
            fmt.Println("3rd Line")
            fmt.Println("4th Line")
        }
    ```
- The above example is pretty simple
    - The "1st Line" message will be the last one printed
- Now a more "complex" example to understand multiples defers
    - With the help of functions
-   ```go
        package main
        
        import "fmt"
        
        func defer01() {
            /* The 'defer fmt.Println("1st Line")'
             * will be deferred
             * That means, that line will be
             * the last one printing a message
             * Because of the 'defer' at the start of this line
             */
            defer fmt.Println("1st Line")
        
            fmt.Println("2nd Line")
            fmt.Println("3rd Line")
            fmt.Println("4th Line")
        }
        
        func defer02()  {
            /* What if we defer more than 1 line?
             * Pretty simple, just need to understand this:
             * "The first will be the last"
             * So the "first defer", will be last executed of the "defers"
             * and the "last defer", will be first executed of the "defers"
            */
            defer fmt.Println("1st Line - 1st Defer")
            defer fmt.Println("2nd Line - 2nd Defer")
            defer fmt.Println("3rd Line - 3rd Defer")
            defer fmt.Println("4th Line - 4th Defer")
        }
        
        func main()  {
            defer01()
            fmt.Println("------------------------")
            defer02()
        }
    ```