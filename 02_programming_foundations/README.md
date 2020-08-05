#### 02 Programming Foundations
1. [_Boolean_](https://golang.org/ref/spec#Boolean_types)
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
2. [_String_](https://golang.org/ref/spec#String_literals)
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
3. [_Constants_](https://golang.org/ref/spec#Constant_declarations)
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
4. [_Iota_](https://golang.org/ref/spec#Iota)
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
07. [Loop _For_](https://golang.org/ref/spec#For_statements)
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
    - Just a _disclaimer_, when doing this:
        -   ```
                for j := 0; j < 5; j++ {
                    fmt.Println("- ", j) // J exists
                }
                fmt.Println(j) // J doesn't exist
            ```
        - The __j__ variable won't be accessible out of the [scope]() of _for loop_
08. [_Conditional_](https://golang.org/ref/spec#If_statements)
    - Conditionals are decisions inside the code
    -   ```go
            package main
            
            import "fmt"
            
            func main() {
            	var numType string
            	for i := 1; i < 10; i++ {
            		if i % 2 == 0 {
            			numType = "even"
            		} else {
            			numType = "odd"
            		}
            		fmt.Println(i, " - ", numType)
            	}
            }
        ```
    -   ```
            package main
            
            import "fmt"
            
            func main() {
            	var numType string
            	for i := 1; i < 10; i++ {
            		if i%2 == 0 {
            			numType = "even"
            		} else {
            			numType = "odd"
            		}
            	}
            
            	switch numType {
            		case "even": {
            			fmt.Println("The last number was an even number!")
            		} 
            		case "odd": {
            			fmt.Println("The last number was an odd number!")
            		}
            		default: {
            			fmt.Println("The values doesn't match any one of the possibilities")
            		}
            	}
            }
        ```
    - The [switch](https://golang.org/ref/spec#Switch_statements) can be adapted to perform certain action depending on the variable type using _interfaces_
        - ```
              package main
              
              import (
                  "fmt"
              )
              
              var unknown interface{}
              
              func main() {
                    switch unknown.(type) {
                    case int:
                        fmt.Println("This is a integer!")
                    case bool:
                        fmt.Println("This is a boolean")
                    case string:
                        fmt.Println("This is a string")
                    case float32, float64:
                        fmt.Println("This is a float!")
                    default:
                        fmt.Println("Unknown type!")
                    }
              }
          ```
09. [_Logical Operator_](https://golang.org/ref/spec#Logical_operators)
    - There are 3 logical operators
        - __&&__ (AND)
            - A _&&_ B = _"if A then B, else false"_
            - |    A|    B|A _&&_ B|
              |-----|-----|--------|
              |false|false|   false|
              |false|true |   false|
              |true |false|   false|
              |true |true |    true|
          
        - __||__ (OR)
            - A _||_ B = _"if A then true else B"_
            - |    A|    B|A _\\\\_ B|
              |-----|-----|--------|
              |false|false|   false|
              |false|true |    true|
              |true |false|    true|
              |true |true |    true|
        - __!__ (NOT)
            - _!_ A = _"not A"_
            - |A    |   !A|
              |-----|-----|
              |false| true|
              |true |false|
    - __DISCLAIMER__
        - The function ``strconv.FormatBool`` receive a boolean as parameter and return it in string form
        - For example _strconv.FormatBool(true)_ will return _"true"_
    -   ```
            package main
            
            import (
            	"fmt"
            	"strconv"
            )
            
            /*
            * In this program I'll be developing
            * a program that will build the same
            * tables shown at README.md file
            */
            func main()  {
            	var values [2]bool
            	values[0] = false
            	values[1] = true
            	/*
            	 * values = [false, true]
            	 */
            	fmt.Println("---And Operator---")
            	for _, valueA := range values {
            		for _, valueB := range values {
            			fmt.Printf("%s && %s = %s\n",
            				strconv.FormatBool(valueA),
            				strconv.FormatBool(valueB),
            				strconv.FormatBool(valueA && valueB))
            		}
            	}
            	fmt.Println("---Or Operator---")
            	for _, valueA := range values {
            		for _, valueB := range values {
            			fmt.Printf("%s || %s = %s\n",
            				strconv.FormatBool(valueA),
            				strconv.FormatBool(valueB),
            				strconv.FormatBool(valueA || valueB))
            		}
            	}
            	fmt.Println("---Not Operator---")
            	for _, valueA := range values {
            		fmt.Printf("!%s = %s\n",
            			strconv.FormatBool(valueA),
            			strconv.FormatBool(!valueA))
            	}
            }
        ```