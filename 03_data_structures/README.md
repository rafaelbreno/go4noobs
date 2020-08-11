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
04. [_Slice Make_](https://golang.org/doc/effective_go.html#allocation_make)
    - [_ref/spec_](https://golang.org/ref/spec#Making_slices_maps_and_channels)
    - The function __make__ takes a _Type T_
        - _Slice_ - The one we're seeing
        - [_Map_]() - _Don't mind this one, we'll see it later_
        - [_Channel_]() - _Don't mind this one, we'll see it later_
    - _make(Type T, length, capacity)_
        - _Type T_ - The variable type it will be creating
        - _Length_ - Initialize the quantity of zeroed values
        - _Capacity_ - Max quantity of elements
    -   ```go
            package main
            
            import "fmt"
            
            func main() {
                /*
                 * Slice of type: []int
                 * length of 5
                 * capacity of 10
                 */
            	slice := make([]int, 5, 10)
            
            	for key, value := range slice {
            		fmt.Println(key, "-", value)
            	}
            }
        ```
        - This will output __5__ _zeroed_ values
    - How was said before the slices are _flexible_
    - So even when the capacity is "enabled", with append it's possible to add more item
        -   ```go
                package main
                
                import "fmt"
                
                func main() {
                
                	fmt.Println("----Len < Cap")
                	slice := make([]int, 3, 5)
		            // 3 of 5
                	fmt.Println("----Overflowing Cap----")
                	for i := 1; i <= 5; i++ {
                		slice = append(slice, i)
                	}
                    //10 out of "5"
                	for key, value := range slice {
                		fmt.Println(key, "-", value)
                	}
                	/*
                	 * When appending, you can overflown the capacity
                	 * because slices are flexible and expandable
                	*/
                }
            ```
    - Reminder: 
        - Multidimensional Slices work a like as multidimensional arrays(seen before at _item 1_)
        -   ```go
                package main
                
                import "fmt"
                
                func main() {
                	fmt.Println("----Multidimensional----")
                	/* This is a 3 dimensional array */
                	var multid [][][]int
                	/* That means every 'i', have and 'j' of 'k'
                	 * If you're already familiarized with math you know about R³(or R3)
                	 * This means there will be a "point" with 3 coordinates (i, j, k)
                	 * and this point will have a value
                	 */
                
                	/* The example below we'll be building
                	 * a slice with 3 "cardinal points"
                	 * (i, j, k) and the value will be the sum of 3 of them
                	 * f(i, j, k) = i + j + k
                	**/
                
                	/* Here it will create an 'i' point
                	 * and assign a slice [][]int{} value to it
                	 */
                	for i := 0; i < 3; i++ {
                		multid = append(multid, [][]int{})
                
                		/* Here it will create a 'j' point
                		 * and assign a slice []int{} value to it
                		 */
                		for j := 0; j < 3; j++ {
                			multid[i] = append(multid[i], []int{})
                
                			/* Here will create the 'k' point
                			 * and assign a integer value to it
                			**/
                			for k := 0; k < 3; k++ {
                				multid[i][j] = append(multid[i][j], i+j+k)
                			}
                		}
                	}
                
                	// Accessing i slice([][]int{})
                	for i_pos, i := range multid {
                		// Accessing j slice([]int{})
                		for j_pos, j := range i {
                			// Accessing k value(int)
                			for k_pos, k := range j {
                				// k_pos is the position of value k
                				fmt.Printf("(%d, %d, %d) = %d\n", i_pos, j_pos, k_pos, k)
                			}
                		}
                	}
                }
            ```
    - _Another_ little __disclaimer__
        - _"If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array."_
        - This is about _"append()"_
        - For example, if create a slice:
            - ```slice := make([]int{}, 3, 5)```
            - Slice of array int
            - 3 zeroed values
            - Capacity of 5
        - Then for some reason need to add 4 more number
            - ``` 
                    addons := []int{1, 2, 3, 4}
                    slice = append(slice, addons...) 
              ```
            - This __will__ exceed the max capacity of 5 by 2
        - This won't throw any error __but__ instead it will create a new _slice_, assign the values and discard the old ones
        - It will bring a more  execution type
    - An example:
        -   ```go
              package main
              
              import (
                "fmt"
                "time"
              )
              
              func main() {
                // Case 1:
                start1 := time.Now()
                slice1 := make([]int, 3, 5)
                slice1 = append(slice1, 1, 2, 3, 4)
                fmt.Println("Time with low cap: ", time.Since(start1))
                
                // Case 2:
                start2 := time.Now()
                slice2 := make([]int, 3, 10)
                slice2 = append(slice2, 1, 2, 3, 4)
                fmt.Println("Time with high cap: ", time.Since(start2))
              
              }
            ```
    - In the code above you will be able to see the difference in time execution in both of them
        - You can totally ignore the _"time"_ package
    - To try to explain what is going on, I will bring it to a RL scenario
    - Our slices will be _egg cartoons_, and the numbers will be, obviously, the _eggs_
    - The program will be a worker that pack those eggs
    - Case 1:
        - The worker receives an egg carton, with 5 slots, and 3 eggs
        - Then he receives and order:
        - _Add 4 more eggs_
        - The _egg cartoon_ has 5 slots, and 7 eggs
        - Because there isn't any more slots
        - The worker will be obtaining another _egg cartoon_ with more slots
        - Then he will remove all eggs, transfer all of them into the new _egg cartoon_
        - After all of this, he will be discarding the old package
    - Case 2:
        - The worker receives an egg carton, with 10 slots, and 3 eggs
        - Then he receives and order:
        - _Add 4 more eggs_
        - The _egg cartoon_ has 10 slots, and 7 eggs
        - Everything is fine
    - You can see that the _Case 1_ have way more steps than _Case 2_
    - That can explain why the difference between those times
    - Allowing you to be more concerned about execution time
05. [_Maps_](https://golang.org/doc/effective_go.html#maps)
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