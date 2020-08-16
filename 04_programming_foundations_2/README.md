#### Programming Foundation 2
1. [_Simple Function_](https://golang.org/doc/effective_go.html#functions)
    - [_Functions_](https://golang.org/ref/spec#Function_types) are an abstracted part of your code
    - It's like the function at math, you send a value(called parameters), then it'll return a value
        - Not always it'll return a value
    - Basic structure:
        -   ```
                func funcName(paramName paramType) returnType {
                    ...func body...
                }
            ```
    - Remember [_Scopes_](https://github.com/rafaelbreno/go4noobs/tree/master/01_variables_and_types)? 
        - No? Go back there and give it a little whirl before continuing here.
    - Right, the _"params"_ and variables declared inside the function can only be accessed _inside_ the function
        - The _function_ body is it's the _function_ scope
    -   ```go
            package main
            
            import "fmt"
            
            /*
             * remember the shown function structure
             * in this case:
             *    - funcName = sum
             *    - paramName(paramType) = num1(int) and num2(int)
             *    - returnType = int
             */
            func sum(num1 int, num2 int) int {
            	/* This is the function scope
            	 * the only things that can be accessed here is:
            	 *     - num1
            	 *     - num2
            	 *     - global variables/constants
            	*/
            
            	// The return type is int, just the sum of the 2 numbers
            	return num1 + num2
            }
            
            func main()  {
            	fmt.Println("5 + 9 =", sum(5, 9))
            	fmt.Println("5 - 9 =", sum(5, -9))
            	fmt.Println("10004638 + 128721671 =", sum(10004638, 128721671))
            	//This won't work because it would be passing float values, not integer
            	//fmt.Println("5.001 + 9.999 =", sum(5.001, 9.999))
            }
        ```
02. [_Variadic Functions_](https://golang.org/ref/spec#Function_types)
    - Variadic functions are functions that can receive _n_ params
    - The responsible for this is the operator __"..."__, and we'll see it a lot of times
    -   ```go
            package main
            
            import "fmt"
            
            // The same example, but now a variadic version
            func sum(nums ...int) int {
            	/* It will form a slice of int
            	*/
            	fmt.Printf("Type: %T\nItems: %d\n", nums, nums)
            	sum := 0
            	for _, num := range nums {
            		sum += num
            	}
            	return sum
            }
            
            func main()  {
            	/* Sending the values
            	 * If I wanted I could add even more number
            	 * a finite amount like, 10, 40, 100, 1000 values
            	 * and just that little function there
            	 * would handle everything
            	*/
            	result := sum(1, 2, 3, 4, 5, 6)
            
            	fmt.Println("\n\n1 + 2 + 3 + 4 + 5 + 6 = ", result)
            }
        ```
03. [_Multiple Return_](https://golang.org/doc/effective_go.html#multiple-returns)
    - In the last two functions, we only returned 1 value
    - Sometimes we want to return 2 values
        - At the moment I'm writing this I'm kinda aware of the "errors"
        - So I'll probably write almost the same thing later, but like this:
        - _"Most of the times you should return +2 values, 1 of them being an __'Error'__ type"_
        - For now, you don't have to care that much about this topic
        - Later we'll be seeing it anyway
    - The principal difference will be at the:
        - _returnType_
        - _return variable_
    -   ```go
            package main
            
            import "fmt"
            
            func multipleReturn(num int, words ...string) (int, string)  {
            	/* The first difference is at returnType
            	 * Before we've seen only 'string', or 'int'
            	 * In this case, we can add n-types inside the ()
            	*/
            	/* This function will be:
            	 * Factorial
            	 * String concatenate
            	*/
            	fac := 1
            	for i := 1; i <= num; i++ {
            		fac *= i
            	}
            	phrase := ""
            	for _, word := range words {
            		phrase += word + " "
            	}
            
            	/* Here is another difference
            	 * like the returnType, we'll be return
            	 * the values separated by a comma ","
            	 * Remember:
            	 * FOLLOWING THE SAME ORDER OF returnType
            	*/
            	return fac, phrase
            }
            
            func main()  {
            //  (int, string) remember the order?
            	factorial, sentence := multipleReturn(8, "Hello!", "This", "is", "a", "multiple", "return", "function")
            
            	fmt.Println(sentence)
            	fmt.Println("Factorial:", factorial)
            }
        ```
04. [_Defer Statement_](https://golang.org/ref/spec#Defer_statements)
    - _"A ['defer'](https://golang.org/doc/effective_go.html#defer) statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body..."_
    - This means that a when you append the statement _defer_ at a line, that line will only be executed/ran at the end of that scope
    -   ```go
            package main
            
            import "fmt"
            
            func main()  {
            	/* The 'defer fmt.Println("1st Line")'
            	 * will be deferred
            	 * That means, that line will be
            	 * the last one printing a message
            	 * Because of the 'defer' at the start of this line
            	 */
            	defer fmt.Println("1st Line")
            
            	fmt.Println("2nd Line")
            	fmt.Println("3rd Line")
            	fmt.Println("4th Line")
            }
        ```
    - The above example is pretty simple
        - The "1st Line" message will be the last one printed
    - Now a more "complex" example to understand multiples defers
        - With the help of functions
    -   ```go
            package main
            
            import "fmt"
            
            func defer01() {
            	/* The 'defer fmt.Println("1st Line")'
            	 * will be deferred
            	 * That means, that line will be
            	 * the last one printing a message
            	 * Because of the 'defer' at the start of this line
            	 */
            	defer fmt.Println("1st Line")
            
            	fmt.Println("2nd Line")
            	fmt.Println("3rd Line")
            	fmt.Println("4th Line")
            }
            
            func defer02()  {
            	/* What if we defer more than 1 line?
            	 * Pretty simple, just need to understand this:
            	 * "The first will be the last"
            	 * So the "first defer", will be last executed of the "defers"
            	 * and the "last defer", will be first executed of the "defers"
            	*/
            	defer fmt.Println("1st Line - 1st Defer")
            	defer fmt.Println("2nd Line - 2nd Defer")
            	defer fmt.Println("3rd Line - 3rd Defer")
            	defer fmt.Println("4th Line - 4th Defer")
            }
            
            func main()  {
            	defer01()
            	fmt.Println("------------------------")
            	defer02()
            }
        ```
05. [_Methods_](https://golang.org/ref/spec#Method_declarations)
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
06. [_Interfaces_](https://golang.org/ref/spec#Interface_types)
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