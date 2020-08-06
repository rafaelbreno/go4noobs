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