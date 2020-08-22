#### Boolean
- [_Doc_](https://golang.org/ref/spec#Boolean_types)
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