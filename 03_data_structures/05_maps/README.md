#### Maps
- [Doc](https://golang.org/doc/effective_go.html#maps)
- _"\[...\] data structure that associate values of one type (the key) with values of another type (the element or value)."_
- That means we a _"key"_ associated to _"value"_ of different types
-   ```go
        package main
        
        import "fmt"
        
        func main()  {
            /* Just explaining the map structure
             * map[key]value
             * the 'key' is what is searched
             * 'value' is what is returned
            **/
            intString := make(map[int]string, 10)
        
            // 02_programming_foundations/02_string
            hello := []byte("hello!!")
        
            for key, value := range hello {
                intString[key] = string(value)
            }
            /* The output will follow the structure
             * key:value , almost like JSON
             * [0:h 1:e 2:l 3:l 4:0 5:! 6:!]
            **/
            fmt.Println(intString)
            
            /* Oh now, I made a little typo
             * By mistake I've added a two "!" instead of one
             * so to remove the last "!", exists a simple function
             * delete(map, key)
             * remember, key is before the ":"
             * e.g: 1:e
             * "1" is the key, "e" is the value
            */
            delete(intString, 6)
    
            fmt.Println(intString)
            // prints [0:h 1:e 2:l 3:l 4:0 5:!]    
        }
    ```
    - [02_programming_foundations/02_string](https://github.com/rafaelbreno/go4noobs/tree/master/02_programming_foundations#02-programming-foundations)
- [_Comma Ok_](https://golang.org/doc/effective_go.html#maps)
    -   ```go
            package main
            
            import "fmt"
            
            func main() {
                /* Just explaining the map structure
                 * map[key]value
                 * the 'key' is what is searched
                 * 'value' is what is returned
                **/
                intString := make(map[int]string, 10)
            
                // 02_programming_foundations/02_string
                hello := []byte("hello")
            
                for key, value := range hello {
                    intString[key] = string(value)
                }
            
                /*
                 * The following verification is called "comma ok"
                 * When you try to get a value from an non-existent key
                 * it will return you a "zeroed" value
                 * To show you if the value is a "0" or zeroed value
                 * "ok" will return a boolean value
                 * true - the value exists
                 * false - value non-existent and zeroed
                **/
                value, ok := intString[13]
                fmt.Println("Value:", value, "- Exists:", ok)
            }
        ```
- Example, mapping functions:
    -   ```go
            package main
            
            import "fmt"
            
            /* Don't need to be scared
             * TODO: Add chapter url
             * We'll be studying function at Chapter
             * func funcName(params) returnType {
             * 		func scope
             * }
            **/
            
            /* Name: returnTrue
             * Params: nothing
             * Return Type: String
             * Return Value: "true"
            */
            func returnTrue() string  {
                return "True"
            }
            
            /* Name: returnFalse
             * Params: nothing
             * Return Type: String
             * Return Value: "false"
             */
            func returnFalse() string  {
                return "False"
            }
            
            func main() {
                /* Here's a little more "advanced" example
                 * we'll be using functions
                 * Don't need to be scared, we'll be entering in that chapter
                 * Here:
                 * FOR NOW You only need to know that a function, is a structure
                 * that returns something
                **/
                funcs := map[int]func()string{
                  //key: function
                    0: returnFalse,
                    1: returnTrue,
                }
            
                // key, value := range map
                for _, f := range funcs {
                    /* r will be the return valued from a function
                     * Again, don't need to get frighten about function
                     * we'll be studying it more deeply again at chapter
                     * TODO: add function chapter link
                    **/
                    r := f()
                    fmt.Println(r)
                }
            }
        ```
    - [Functions]()
- __Disclaimer!!__
    - Maps don't have an order, so every _loop_ can be different in a _different order_