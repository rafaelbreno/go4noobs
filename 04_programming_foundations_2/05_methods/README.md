#### Methods
- [Doc](https://golang.org/ref/spec#Method_declarations)
- _"A [method](A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.) is a function with a __receiver__. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type."_
- Here we'll be re-studying the function structure
-   ```
        func funcReceiver funcName(paramName paramType) returnType { funcScope }
    ```
- At the function section I didn't have written about the func _receiver_ just to no worry about at the time
- Now we'll be seeing it
    - If you've already studied about _object oriented programming_ this section will be easier(and kinda controversial)
-   ```go
        package main
        
        import "fmt"
        
        /* Creating a simple struct
         * will have 2 arguments
         * nameParts and completeName
         * completeName will be the concatenation of nameParts
        */
        type person struct {
            nameParts []string
            completeName string
        }
        /* Here we'll be receiving a Pointer
         * Thinks that a pointer, points to something
         * the case bellow
         * 'p' points to a variable of type 'person'
        */
        func (p *person) concatName()  {
            /* Because 'p' is a pointer of 'person'
             * we can access and modify it's values
             * without the need of returning something
            */
            for _, namePart := range p.nameParts {
                // Accessing and modifying completeName
                p.completeName += namePart + " "
            }
        }
        
        func main()  {
            // Instantiating a variable 'john' of type 'person'
            john := person{
                nameParts:    []string{"John", "Foo", "Doe", "Bar"},
                completeName: "",
            }
            /* In this case 'john' will be receiver
             * 'concatName' will receive 'john' pointer
            */
            john.concatName()
        
            fmt.Println(john.nameParts)
            fmt.Println(john.completeName)
        }
    ```
- This kinda remembers the "methods" from OOP
    - In the place of a _"class"_ we've the struct type
    - In the place of a _"object"_ we've the variable of that struct
    - In the place of a _"method"_ we've the function with a _"receiver"_ specifically for that struct