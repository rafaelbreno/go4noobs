#### 02 Programming Foundations
1. Boolean
    - Boolean is a binary type(can assume true[1] or false[0])
    - Its value can be determined by relational operators(*"==", "!=", ">", "=>"* are some examples)
        - *"=="* - Equal ...
        - *"!="* - Different ...
        - *">"*  - Greater than ...
        - *">="* - Greater or equal than ...
    - ```
          func main() {
            a := 10
            b := 20
            fmt.Printf("a = %v |||  b = %v\n", a, b)
            fmt.Printf("a == b : %v\n", a == b)
            fmt.Printf("a != b : %v\n", a != b)
            fmt.Printf("a >= b : %v\n", a >= b)
            fmt.Printf("a <= b : %v\n", a <= b)
          }
      ```
2. String
    - String is an array os _characters_
    - ```go
        package main
        import "fmt"
        func main()  {
        	s := "Hello, World!"
        	sb := []byte(s)
        
        	for _, v := range sb {
        		fmt.Printf("%v - %T <-> %#U - %#x\n", v, v, v, v)
        	}
        
        	// Char per char
        	for _, z := range s {
        		fmt.Printf("%v - %T <-> %#U - %#x\n", z, z, z, z)
        	}
        
        	// Byte per Byte
        	for i := 0; i < len(s); i++ {
        		fmt.Printf("%v - %T <-> %#U - %#x\n", s[i], s[i], s[i], s[i])
        	}
        }
      ```
    - The _%v_ is already known, but what's the _#_ in the middle of _%#U_ and _%#x_
        - __#__ is known as _Flag_, it's 1 in 5
    - _[Flags:](https://golang.org/pkg/fmt/)_
        - _+ :_ Always print a sign for numeric values
        - _- :_ Pads with space rather on left than right
        - _# :_ Alternate the format
        - _' '(space):_ Leave a space for elided sign on numbers _(% b)_
        - _0 :_ Pad with leading zeros
    - [More about string](https://blog.golang.org/strings)
3. Constants
    - Const, located outside the scope of functions
    - The assigned value doesn't imply the type
    - E.g
        - _const c = 10_
        - __c__ can be: int, float, string, etc.
        - It will only assume a type when a variable calls it
    - ```go
        package main
        
        import (
            "fmt"
        )
      
        const c = 10
      
        var f float32
      
        var i int
      
        func main() {
            f = c
            i = c

            fmt.Println(f)
            fmt.Println(i)
        }
      ```
    - _Const_ type, defined when used 
    - _Var_ type, defined when assigned
    - Const can also be declared like this:
        - ```
          const {
              a = 10
              b = 20
              c = 30
          }
          ```
4. Iota
    - Iota by _first sight_ looks like a counter inside const's scope
        - ```
            const (
                a = iota // iota == 0
                b = iota // iota == 1
                c = iota // iota == 2
                d = iota // iota == 3
            )
          ```
05. Part 1 - Exercises
    1. Problem 01
        - Write a program that shows a number in:
            - Decimal - Base 10
            - Binary - Base 2
            - Hexadecimal - base 16
        - ```go
          package main
          
          import "fmt"
          
          func main() {
            var num int
            // Decimal - Base 10
            fmt.Printf("%d", num)
            // Binary - Base 2
            fmt.Printf("%b", num)
            // Hexadecimal - Base 16
            fmt.Printf("%#x", num)
          }
          ```
06. Procedural Programming
    - Procedural programming is a programming paradigm
        - The simplest way to explain it is: the program will be executed line by line
    -   ```go
            func main() {
            	fmt.Printf(var01) // Flag 01
            	var01 := "Hello Coder!!" // Flag 02
            	fmt.Printf(var01) // Flag 03
            }
        ```
    - If you run the code above will get an error, why?
    - Because you printed the variable before declaring it
        - At the _Flag 01_ the program don't "know" the existence of the variable __var01__ _yet_
            - Here the program will brake, but just for study purposes we will continue the logical processes
        - Then at the _Flag 02_ where finally the __var01__ were declared the program will know the existence of it
        - Finally at _Flag 03_ the program will be able to print __var01's__ value
    - __Disclaimer!!!__ This is the basic explanation about procedural paradigm
    - __I truly__ recommend you reserve a tiny bit of your day to give it a try.
    - Just to search about some "rules" and "good behavior" when programming in _Go_
07. Loop _For_
    - In go every loop is a _for loop_
        - Do
        - Do While
        - Foreach
    - Why?
        - Simple, because every loop can be constructed using simple for
    -   ```go
            package main
            
            import "fmt"
            
            func main()  {
            	// Types of for
            	var i int
            	var n []int
            	fmt.Print("\n------|Do While i < 3|------\n")
            	for i < 3 {
            		fmt.Println("Loop 1: ", i)
            		i = i + 1
            	}
            	fmt.Print("\n------|Common For|------\n")
            	for i = 0; i < 5; i++ {
            		fmt.Println("Loop 2: ", i)
            	}
            	i = 0
            	// Do While i != 3
            	fmt.Print("\n------|Do While i != 3|------\n")
            	for {
            		if i == 3 {
            			break
            		}
            		fmt.Println("Loop 3: ", i)
            		i++
            	}
            	fmt.Print("\n------|Only shows n + 1|------\n")
            	i = 0
            	for i < 6 {
            		i++
            		if i % 2 == 0 {
            			continue
            		}
            		fmt.Println("Loop 4: ", i)
            	}
            
            	for i = 1; i < 6; i++ {
            		n = append(n, i)
            	}
            
            	fmt.Print("\n------|Foreach For|------\n")
            	for i = 0; i < len(n); i++ {
            		fmt.Println("Foreach 1: ", n[i])
            	}
            	fmt.Print("\n------|Foreach Range|------\n")
            	/*
            	 * key, value := range arr
                 * [ 
                 *      key1: value1, 
                 *      key2: value2,
                 *      ... ,
                 *      keyN: valueN 
                 * ]
            	*/
            	for key, value := range n  {
            		fmt.Printf("%x: %x\n", key, value)
            	}
            }
        ```