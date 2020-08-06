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
            	/*
            	 * how that works?
            	 * append is a function that receive 2 params
            	 * append(slice, element) = [slice, element]
            	 * so, to create a prepend just invert the order
            	 * append(element, slice), but in this case, element isn't a slice type, and slice isn't a element
            	 * to convert them we do this
            	 * append([]int{element}, slice...)
            	 * "slice..." will transform slice into element
            	 * "[]int{element}" will transform element in a slice of int with unique item element
            	 * kinda confusing but that makes sense
            	*/
            }
        ```