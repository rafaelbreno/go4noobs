#### Functions
- [Doc](https://golang.org/doc/effective_go.html#functions)
- [_Functions_](https://golang.org/ref/spec#Function_types) are an abstracted part of your code
- It's like the function at math, you send a value(called parameters), then it'll return a value
    - Not always it'll return a value
- Basic structure:
    -   ```
            func funcName(paramName paramType) returnType {
                ...func body...
            }
        ```
- Remember [_Scopes_](https://github.com/rafaelbreno/go4noobs/tree/master/01_variables_and_types)? 
    - No? Go back there and give it a little whirl before continuing here.
- Right, the _"params"_ and variables declared inside the function can only be accessed _inside_ the function
    - The _function_ body is it's the _function_ scope
-   ```go
        package main
        
        import "fmt"
        
        /*
         * remember the shown function structure
         * in this case:
         *    - funcName = sum
         *    - paramName(paramType) = num1(int) and num2(int)
         *    - returnType = int
         */
        func sum(num1 int, num2 int) int {
            /* This is the function scope
             * the only things that can be accessed here is:
             *     - num1
             *     - num2
             *     - global variables/constants
            */
        
            // The return type is int, just the sum of the 2 numbers
            return num1 + num2
        }
        
        func main()  {
            fmt.Println("5 + 9 =", sum(5, 9))
            fmt.Println("5 - 9 =", sum(5, -9))
            fmt.Println("10004638 + 128721671 =", sum(10004638, 128721671))
            //This won't work because it would be passing float values, not integer
            //fmt.Println("5.001 + 9.999 =", sum(5.001, 9.999))
        }
    ```