#### Closure
- Links
    - [Go Tour - Closure](https://tour.golang.org/moretypes/25)
    - [Go by Example - Closure](https://gobyexample.com/closures)
- Closure is a technique to implement a function scope into _something_
- In this case, this _"something"_ will be a variable that will assume a function
-   ```go
        package main
        
        import "fmt"
        
        func main()  {
        
            /* This will instantiate a new
             * function scope */
            a := clos(10)
        
            fmt.Println(a()) // 11
            fmt.Println(a()) // 12
            fmt.Println(a()) // 13
        
            /* This will instantiate a new
             * function scope */
            b := clos(245)
        
            fmt.Println(b()) // 246
            fmt.Println(b()) // 247
            fmt.Println(b()) // 248
        }
        
        func clos(i int) func() int {
            x := i
            return func() int {
                x++
                return x
            }
        }
    ```