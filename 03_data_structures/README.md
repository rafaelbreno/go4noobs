#### 03 Data Structures
1. [Array](https://golang.org/ref/spec#Array_types)
    - Array is a data structure, that is formed with multiples values
    - For example, a _string_ can be called _array of chars_
    -   ```
            package main
            
            import "fmt"
            
            var arr1 [5]int
            var arr2 [5][5]int
            func main()  {
            	fmt.Println("----Array 1 Dimension-----")
            	for i := 0; i < 5; i++ {
            		arr1[i] = i * 10
            	}
            	for key, value := range arr1 {
            		println(key, " => ", value)
            	}
            	fmt.Println("\n----Array 2 Dimensions-----")
            	for i := 0; i < 5; i++ {
            		for j := 0; j < 5; j++ {
            			// "+ 1" because the first value is 0
            			arr2[i][j] = (i + 1) * (j + 1)
            		}
            	}
            	/*
            	arr2 = [
            		[1, 2, 3, 4, 5],
            		[2, 4, 6, 8, 10],
            		[3, 6, 9, 12, 15],
            		[ ... ]
            	]
            	*/
            	for _, valuei := range arr2 {
            		/*
            			valuei = [1, 2, 3, 4, 5]
            			valuei = [2, 4, 6, 8, 10]
            			valuei = [3, 6, 9, 12, 15]
            			....
            		*/
            		for _, valuej := range valuei {
            			/*
            				valuej = 1
            				valuej = 2
            				valuej = 3
            				valuej = 4
            				....
            			*/
            			fmt.Print(valuej, "\t")
            		}
            		fmt.Print("\n")
            	}
            }
        ```
    - Reviewing the assignment of array
        - `var {name} [{length}]type`
        - __name__: is the variable name
        - __length__: the quantity of items that this array can assign
        - __type__: the array's items type(int, bool, string, _etc_)
    - In _Go_ arrays are the base of _Slices_(next item)
    - [__Here__](https://golang.org/doc/effective_go.html#arrays) you can read the official documentation about Go's array
2. [_Slices_](https://golang.org/ref/spec#Slice_types)
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
3. [_Slice of Slices_](https://golang.org/doc/effective_go.html#slices)
    - To deal with a Slice of a Slice, it's need to use the following formats:
        - ``slice = slice[:]``
        - ``slice = slice[a:]``
        - ``slice = slice[:b]``
        - ``slice = slice[a:b]``
    - The most _"educational"_ way to explain what the _"a"_ and _"b"_ is, is with numerical sets/intervals
        - _"a"_ is left-closed
        - _"b"_ is right-open
        - For example, we have a numerical set/interval R = [1, 10]
            - R = 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 
        - If I asked you to select the interval from 3 to 8, you would do something like this
            - _[3, 8]_
            - _right-close_ and _left-closed_
        - _But_ if I asked you to select the interval from 3 to 8, but _right-open_ and _left-closed_, you would need to do this:
            - _[3, 9[_
            - The __"9["__ won't include _"9"_
    -   ```go
            package main
            
            import "fmt"
            
            //                 0  1  2  3  4  5  6
            var slice = []string{"a", "b", "c", "d", "e", "f", "g"}
            
            func main()  {
            	fmt.Println("-----[:]-----")
            	slice1 := slice[:]
            	for key, value := range slice1 {
            		fmt.Println(key, ":", value)
            	}
            	fmt.Println("\n-----[3:]-----")
            	slice2 := slice[3:]
            	for key, value := range slice2 {
            		fmt.Println(key, ":", value)
            	}
            	fmt.Println("\n-----[:3]-----")
            	slice3 := slice[:3]
            	for key, value := range slice3 {
            		fmt.Println(key, ":", value)
            	}
            	fmt.Println("\n-----[2:5]-----")
            	slice4 := slice[2:5]
            	for key, value := range slice4 {
            		fmt.Println(key, ":", value)
            	}
            
            	fmt.Println("\n-----[:2] U [4:]-----")
            	/*
            	 * Again with math notations
            	 * remember "U"?
            	 * That means "union"
            	 * In this case [:2] U [4:]
            	 * Will be uniting the terms from 0 to 1, and from 4 until the last one
            	 * To unite them both we'll be using append with the "..." operator AGAIN!
            	*/
            	union := append(slice[:2], slice[4:]...)
            	fmt.Println(union)
            }
        ```