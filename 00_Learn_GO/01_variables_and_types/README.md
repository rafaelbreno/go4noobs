#### 01 Variables and Types
- This section is an introduction to GO
1. Hello World
    - Like all tutorials we need to make a *Hello World* program
    - In this example we can see the structure of a basic GO program
    - ```go
          package main
          
          import (
            "fmt"
          )
          
          func main() {
            fmt.Println("Hello World,", "I'm Rafael Breno!")
          }
      ```
    - When we run it with:
        - > go build main.go
        - output: "Hello World, I'm Rafael Breno!"
2. Println()
    - I think it's a good habit read the [Official Documentation](https://golang.org/doc/) of any lang/framework/tool that you're using
    - This case is a good example of that
    - Here is the [Println() Func Doc](https://golang.org/pkg/fmt/#Println), reading its possible to do this:
    - ```go
         func main() {
            //Here is  simple function
            bytes, err := fmt.Println("Hello World,", "I'm Rafael Breno!")
            fmt.Println("The string has", bytes, "bytes. Error:", err)
         }
      ```
    - How do I know what's returning, and in what order? Following the doc:
        - "*It returns the number of bytes written and any write error encountered.*"
3. Gopher
    - Gopher is a short operator Its symbol is *":="* 
    - Different of *"="*, that is a *assignment operator*
    - You can see more of it in these links:
        - [Official Doc - Short Variable Declarations](https://golang.org/ref/spec#Short_variable_declarations)
        - [Unofficial Doc - Stackoverflow Question](https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go/45654233#45654233)
    - At the Stackoverflown question, I've founded these *"rules"* about the **":="**
        1. You can't use *":="*  out of *"funcs"* . It's because, out of any func, a statement should start with a keyword.
            - ```go
              // no keywords below, illegal.
              illegal := 42
              
              // `var` keyword makes this statement legal.
              var legal = 42
              
              func foo() {
                alsoLegal := 42
                // reason: it's in a func scope.
              }
              ```
        2. You can't use them twice (in the same scope):
            - ```go
              legal := 42
              legal := 42 // <-- error
              ```
        3. You can use them for multi-variable declarations and assignments:
            - ```go
              foo, bar   := 42, 314
              jazz, bazz := 22, 7
              ```
        4. You can use them twice in **"multi-variable"** declarations, if one of the variables is new:
            - ```go
              foo, bar  := someFunc()
              foo, jazz := someFunc()  // <-- jazz is new
              baz, foo  := someFunc()  // <-- baz is new
              ```
        5. You can use the short declaration to declare a variable in a newer scope even that variable is already declared with the same name before:
            - ```go
              var foo int = 34
              
              func some() {
                // because foo here is scoped to some func
                foo := 42  // <-- legal
                foo = 314  // <-- legal
              }
              ```
        6. You can declare the same name in short statement blocks like: **if**, **for**, **switch**:
            - ```go
              foo := 42
              if foo := someFunc(); foo == 314 {
                // foo is scoped to 314 here
                // ...
              }
              // foo is still 42 here
              ```
    - With it said, we can now use the function [Printf]() and the [Printing "verbs"](https://golang.org/pkg/fmt/#hdr-Printing)
    - To help understand that *":="* guess what type of value it's assigned to it
    - ```go
      func main(){
        numInt := 20
        numFloat := 10.2
        textString := "Olá, meu nome é Rafael"
        textChar := 'R'
        fmt.Printf("Var: %b , Type: %T\n", numInt, numInt)
        fmt.Printf("Var: %b , Type: %T\n", numFloat, numFloat)
        fmt.Printf("Var: %b , Type: %T\n", textString, textString)
        fmt.Printf("Var: %b , Type: %T\n", textChar, textChar)
      }
      ```
4. Scope
    - Scope is an area with limits defined by ``{}``
        - For Example``{ ...in scope... } ...out of scope`` 
    - A good example of it is a "Global Variable", its **scope** is the entire code, can be used in every *func* of a package
    - At the code below we can see that the variable **global** can be used 
    - ```go
          var global = 10
          func main() {
            local := 20
            fmt.Print("Global Variable:", global)
            fmt.Print("\nLocal Variable:", local)
            Scope(local)
          }     
          func Scope(x int){
            //This is legal
            fmt.Print("\n\nGlobal Variable:", global)
          
            //This is legal
            fmt.Print("\nParameter Variable:", x)
          
            //This is illegal, will show 'undefined' error in terminal
            // fmt.Print("Main Variable:", local)
          }
      ```
5. Types
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
    - ```go
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
6. Custom Type
    - In Go there is var types that the developer creates
    - This is something usual in the future, for now its more like an abstract idea
    - ```go
        // type myType subType
        type myType int
        var b myType
        func main() {
            a := 10
            fmt.Printf("%T", a)
            fmt.Printf("\n%T", b)
            /**
            * a = b this is wrong
            * because 'a' is type 'int'
            * and b is type 'myType'
            **/
        }
      ```
    - This program should output:
        ```
          int
          main.myType
        ```
    - Even with the same values "*a(int)*" != "*b(myType)*", because of the var type
7. Type Conversion
    - To convert the type, it's just needed to use a simple function
    - ``var a typeA = typeA(b typeB)``
    - ```go
          //Creating a type based on 'int'
          type myType int
          var b myType = 20        
          func main() {
            a := 10
            fmt.Println("Passing value from b to a:")
            fmt.Printf("b myType -> Type %T , Value %v\n", b, b)
            a = int(b)
            fmt.Printf("a int -> Type %T , Value %v\n", a, a)
            fmt.Println("Passing value from a to b:")
            b = myType(a)
            fmt.Printf("b myType -> Type %T , Value %v\n", b, b)
            fmt.Printf("a int -> Type %T , Value %v\n", a, a)
          }
      ```
##### Exercises
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