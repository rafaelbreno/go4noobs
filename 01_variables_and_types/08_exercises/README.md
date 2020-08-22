#### Exercises
1. 01_example
    - Using the short operator, assign these values with the following identification *"x"*, *"y"* and *"z"*:
        1. 42
        2. "James Bond"
        3. true
    - Now print the values with:
        1. An only Print
        2. Multiples Prints
    - ```go
          func main()  {
            x, y, z := 42, "James Bond", true
            fmt.Printf("x: %v - y: %v - z: %v\n", x, y, z)
            fmt.Printf("x: %v\n", x)
            fmt.Printf("y: %v\n", y)
            fmt.Printf("z: %v", z)
          }
      ```
2. 02_example
    - Use var to declare three variables, they must be at *package-level scope*
    - Don't assign values to them, to identify each:
        1. *"x"* must be type **int**
        2. *"y"* must be type **string**
        3. *"z"* must be type **bool**
    - In *func main*:
        1. Show the values of each one
        2. The compiler assigned values to them, how are they called?
            - They're called *"zero value"*
    - ```go
        var x int
        var y string
        var z bool
        func main() {
            fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)
        }
      ```
3. 03_example
    - Using the solution from previous example
    - In *package-level scope* assign the following values to the variables
        1. Assign **42** to *"x"*
        1. Assign **"James Bond"** to *"y"*
        1. Assign **true** to *"z"*
    - In *func main*
        1. Use *fmt.Sprintf* to assign all these values to a single variable of type string a name *s* using the short operator
        2. Print the variable *s*
    - ```go
          var x int = 42
          var y string = "James Bond"
          var z bool = true
          func main() {
            s := fmt.Sprintf("x: %v\ny: %v\nz: %v\n", x, y, z)
            fmt.Printf("s: %v", s)
          }
      ```
4. 04_example
    - Create a type based on *int*
    - Create a variable to this type, named *"x"*, using the key-word var
    - In *func main*
        1. Show the value of *"x"*
        2. Show the type of *"x"*
        3. Assign 42 to the variable *"x"* with *"="* operator
        4. Show the value of *"x"*
    - ```go
          type myType int
          var a myType
          func main()  {
            fmt.Printf("Value: %v\n", a)
            fmt.Printf("Type: %T\n", a)
            a = 42
            fmt.Printf("Value: %v\n", a)
          }
      ```
5. 05_example
    - Using the solution from previous example
    - In *package-level scope*, using the key-word *"var"*, create a variable named *"y"*
    - The type of this variable must be a custom created type based on *"int"*
    - In *func main*
        1. Show the value of *"x"*
        2. Show the type of *"x"*
        3. Assign 42 to the variable *"x"* with *"="* operator
        4. Show the value of *"x"*
        5 Use the conversion to convert the *"x"* value and assign it to *"y"* var
        1. Show the value of *"y"*
        2. Show the type of *"y"*
    - ```go
        type myType int
        var x myType
        var y int
        func main()  {
            fmt.Printf("Value X: %v\n", x)
            fmt.Printf("Type X: %T\n", x)
            x = 42
            fmt.Printf("Value X: %v\n", x)
            y = int(x)
            fmt.Printf("Value Y: %v\n", y)
            fmt.Printf("Type Y: %T\n", y)
        }
      ```