#### Const
- [Doc](https://golang.org/ref/spec#Constant_declarations)
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