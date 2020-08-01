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
    - Const are assigned outsite the scope of functions
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
     