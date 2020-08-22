#### Interfaces
- [Doc](https://golang.org/ref/spec#Interface_types)
- _"An [_interface_](https://golang.org/doc/effective_go.html#interfaces_and_types) type specifies a method set called its interface. A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to implement the interface. The value of an uninitialized variable of interface type is nil."_
- As the _"method"_, if you have a little knowledge in _OOP_ this step will be easier for you
    - An interface is like a class(_Parent class_) with base methods
    - That other classes(children class) can extend it and inherit those methods
    - That's called _Inheritance_ and _Polymorphism_
- Now back to _Go_
- Here we'll've
    - _Interface_ as Parent class
    - _Struct_ as Children class
-   ```go
        package main
        
        import "fmt"
        
        /************|mainSimple - Start|************/
        
        // Person struct - Parent class
        type Person struct {
            nameParts    []string
            completeName string
            age 		 int
        }
        
        // Firefighter extends Person
        type Firefighter struct {
            Person
            firesFought int
        }
        
        // Teacher extends Person
        type Teacher struct {
            Person
            kidsTaught int
        }
        
        // This functions receives Teacher that extends Person
        func (t Teacher) simpleWork() {
            fmt.Println("I'm ", t.completeName, "and I'm a teacher!")
        }
        
        // This function receives Firefighter that extends Person
        func (f Firefighter) simpleWork() {
            fmt.Println("I'm ", f.completeName, "and I'm a firefighter!")
        }
        
        // This function receives Person
        func (p *Person) concatName() {
            for _, namePart := range p.nameParts {
                p.completeName += namePart + " "
            }
        }
        
        // John is an instance of Teacher that extends Person
        var john = Teacher{
            Person:     Person{
                nameParts:    []string{"John", "Doe"},
                completeName: "",
                age:          21,
            },
            kidsTaught: 110,
        }
        
        // Joanna is an instance of Firefighter that extends Person
        var joanna = Firefighter{
            Person:      Person{
                nameParts:    []string{"Joanna", "Doe"},
                completeName: "",
                age:          26,
            },
            firesFought: 97,
        }
        
        /* The function below will be a simple one
         * with the use of an interface
        */
        func mainSimple()  {
            // Concatenating John Name
            john.concatName()
            // John work
            john.simpleWork()
        
            // Concatenating Joanna Name
            joanna.concatName()
            // Joanna work
            joanna.simpleWork()
        
        }
        /************|mainSimple - End|************/
        
        
        /************|mainComplex - Start|************/
        type Human interface {
            /* Here we've bound the simpleWork function
             * so if there's a 'instance' of Human's interface
             * it'll be accessible
            */
            simpleWork()
        }
        func complexWork(h Human)  {
            //Receiving h type (Teacher of Firefighter)
            switch h.(type) {
                case Teacher:
                    fmt.Println("I'm ", h.(Teacher).completeName, "and I'm a teacher!")
                    break
                case Firefighter:
                    fmt.Println("I'm ", h.(Firefighter).completeName, "and I'm a firefighter!")
                    break
            }
            /* Just an example that it's possible
             * to access simpleWork, and it'll work fine
             * just as in mainSimple
            */
            h.simpleWork()
        }
        
        
        func mainComplex()  {
            complexWork(john)
        
            complexWork(joanna)
        }
        /************|mainComplex - End|************/
        
        func main()  {
            /* This function will be a simple
             * example with 'person' struct
             * that will have a job teacher/firefighter
            **/
            mainSimple()
        
            fmt.Println("------------------------------------")
        
            /* This function will be a more 'complex' one
             * with the same base struct, but
             * now we'll have an interface
            **/
            mainComplex()
        }
    ```