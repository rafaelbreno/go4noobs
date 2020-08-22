#### Slice
- [Doc](https://golang.org/ref/spec#Slice_types)
- More info [__here__](https://golang.org/doc/effective_go.html#slices)
- Slices are a type of _array_ but __waaaay__ more flexible
- The main difference is when you declare it
    - _array:_ `var {name} [{length}]type`
    - _slice:_ `var {name} []type`
- Slice don't have a finite length
-   ```go
        package main
        
        import "fmt"
        
        var slice1 []int
        var slice2 []int
        func main()  {
            /* append will a insert a value at the end of the slice
             * Imagine that a slice is a queue of items
             * appending something will put it at the end of the queue
             */
            fmt.Println("----prepend----")
            for i := 0; i < 5; i++ {
                // This for will append values from 0 to 4
                slice1 = append(slice1, i)
            }
            for _, value := range slice1 {
                // Here will show the items at the order they were appended
                fmt.Println(value)
            }
        
            /*
             * What if I wanted to prepend the values, for some reason
             * "prepend" a value is putting the item at the beginning of the queue
             */
            fmt.Println("----prepend----")
            for i := 0; i < 16; i++ {
                if i % 2 == 0 {
                    // If is a even number
                    slice2 = append(slice2, i)
                } else {
                    // If is a odd number
                    slice2 = append([]int{i}, slice2...)
                }
            }
            fmt.Println(slice2)
        }
    ```
- At the code above you saw the function append()
    - append receive _+2_ params
        - slice
        - element
    - Why _"+2"_?
        - Because append() is a variadic function
- Variadic functions:
    - Is a function that can receive infinite elements
-   ```go
        package main
    
        import "fmt"
    
        var sliceBase []int
        var sliceBase2 []int
        var sliceBase3 []int
        var sliceAdd = []int{1, 2, 3, 4, 5, 6}
        
        func main(){       
            // This:
            fmt.Println("\n-----Base1-----")
            for _, value := range sliceAdd {
                sliceBase = append(sliceBase, value)
            }
            fmt.Println(sliceBase)
        
            // Is the same of this:
            fmt.Println("\n-----Base2-----")
            /*
                Look here
                append([]interface{}, 1, 2, 3, 4, ..., n - 1, n)
            */
            sliceBase2 = append(sliceBase2, 1, 2, 3, 4, 5, 6)
            fmt.Println(sliceBase2)
        
            // And this:
            fmt.Println("\n-----Base3-----")
            sliceBase3 = append(sliceBase3, sliceAdd...)
            fmt.Println(sliceBase3)
        }
    ```
- The code above can explain why when I wanted to _"prepend"_ a value I did this:
    - ``slice = append([]int{1}, slice...)``