#### Slice Make
- [Doc](https://golang.org/doc/effective_go.html#allocation_make)
- [_ref/spec_](https://golang.org/ref/spec#Making_slices_maps_and_channels)
- The function __make__ takes a _Type T_
    - _Slice_ - The one we're seeing
    - [_Map_]() - _Don't mind this one, we'll see it later_
    - [_Channel_]() - _Don't mind this one too, we'll see it later_
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