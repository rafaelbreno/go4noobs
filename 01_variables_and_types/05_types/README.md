#### Types
- The title is self-explanatory, type is the *"format"* of a designated variable
- Assigning a type for a variable it means that variable can only accept it until the end of the program
- These types can be(some examples):
    - Primitive
        - | Title   | Type    | Example   |
          |---------|---------|-----------|
          | Integer | int     | 10        |
          | Float   | float32 | 10.2      |
          | Boolean | bool    | true      |
          | String  | string  | "Example" |
    - [Derived](https://www.tutorialspoint.com/go/go_data_types.htm)
        - Pointer
        - Array
        - Structure
        - Union
        - Function
        - Slice
        - Interface
        - Map
        - Channel
- For now I'm going to stick with only *Primitive* types
```go
var a = 10
var b float32
/*
Can't assign a value after declaring a variable out of a func
b = 20
*/

func main()  {
  b = 20.147
  var c bool = true
  var d string = "Rafael"
  fmt.Printf("Var: %v , Type: %T\n", a, a)
  fmt.Printf("Var: %v , Type: %T\n", b, b)
  fmt.Printf("Var: %v , Type: %T\n", c, c)
  fmt.Printf("Var: %v , Type: %T", d, d)
}
```
