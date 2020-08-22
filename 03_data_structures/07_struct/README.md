#### Struct
- [Doc](https://golang.org/ref/spec#Struct_types)
- _"A struct is a sequence of named elements, called fields, each of which has a name and a type..."_ 
- _Roughly_ saying, a struct is a custom _Type_ formed by other types, native or other _custom types(other structs)_
-   ```go
        // main.go
        package main
        
        import "fmt"
        
        /*
         * For example we can create a struct of a person
         * and it will receive information that someone can have
         * Lie: name, height, adult(if he\she is an adult)
        */
        var person struct {
        //	varName 	varType
            name 		string
            height 		float64
            adult 		bool
        }
        
        func main()  {
            // Let's create Me
            rafael := person
            
            //Defining a value to it
            // structVar.varName = value
        
            // My name is Rafael Breno
            rafael.name = "Rafael Breno"
            // I'm 1.83m tall, YES METERS NOT FEET, don't judge me
            rafael.height = 1.83
            // Yes, I'm an adult
            rafael.adult = true
            fmt.Println("--------")
            fmt.Println("Name:", rafael.name)
            fmt.Println("Height:", rafael.height, "m")
            fmt.Println(rafael.name, "is an adult? ", rafael.adult)
            fmt.Println("--------")
        
            /* Here you will see something that remembers maps
             * but without the key, just the value
            */
            fmt.Println("Struct: ", rafael)
        
            /* When printing the Type of a struct you will se 2 parts
             * struct - showing that it's a struct
             * and it's "internal components", internal structure of other types
            */
            fmt.Printf("Type: %T", rafael)
        }
    ```
- The example above was a simple example of a simple _struct_, the next one we'll me mixing things up
- The next example we'll see a map of a struct that have another struct
-   ```go
        // main02.go
        package main
        
        import "fmt"
        
        // Let's create a struct for dates
        type date struct {
            day, month, year int
        }
        
        /* Like the last example
         * This is a struct of a person
         * and it'll be using the date struct
         * to store the birthday
        */
        type student struct {
            name 		string
            age 		int
        
            /* Here is our custom type
             * the date struct
            */
            birthday 	date
        }
        
        func main()  {
            /* Disclaimer!!
             * Here I've forgot how to create a map
             * So I'd to go back in this documentation
             * to remember how to create one
            */
            students := make(map[int]student, 3)
        
            /* At the map examples, you should remember the "Hello!!" one
             * There I'd a map[int]string
             * string was the value value and int the key
             * Here will be the same
             * but instead of "string"
             * We'll be filling a struct
            */
        
            // Creating an empty student struct
            rafael := student{}
            // Filling the basic "inputs"
            rafael.name = "Rafael Breno"
            rafael.age = 20
        
            // For the date we need to create a var
            // a var of type date
            rafaelBirthDate := date{}
        
            // fill the info of this variable
            rafaelBirthDate.day = 22
            rafaelBirthDate.month = 11
            rafaelBirthDate.year = 1999
        
            // then fill the birthday
            rafael.birthday = rafaelBirthDate
        
            // Adding it to our map
            students[0] = rafael
            /*
             * There's another way to do the same thing above
             * For me it's easier to visualise what's happening
             */
            students[1] = student{
                name:     "Rafael Breno",
                age:      20,
                birthday: date{
                    day:   22,
                    month: 11,
                    year:  1999,
                },
            }
        
        
            fmt.Println("---First Method---")
            fmt.Println(students[0])
            fmt.Println("\n---Second Method---")
            fmt.Println(students[1])
        
            // It's the same thing, done In different ways
        }
    ```