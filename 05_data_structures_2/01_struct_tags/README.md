#### Struct Tags & Reflect Package
- [_Tag_](https://golang.org/ref/spec#Tag)
- [_Reflect_](https://golang.org/pkg/reflect/)
	- Package to iterate a struct
- At [_Data Structures I_](https://github.com/rafaelbreno/go4noobs/tree/master/03_data_structures) we saw a simple example of [_Struct_](https://github.com/rafaelbreno/go4noobs/tree/master/03_data_structures/07_struct)
	- 	```go
			type person struct {
				Name 	string
				Age		int
				Country	string
			}
		```
- Sometimes you should feel a "lack" of info in your struct
- You even can think about adding a struct, slice, arrays, maps inside your main _struct_ to supply this lack of info
- That was my case when searching some info/docs to develop the Request Validator
	- That's where I found out about [_Tags_](https://golang.org/ref/spec#Tag)
	- It's a way to add more info in one "field" of a struct
- 	```
		package main

		import (
			"fmt"
		// Importing Reflect package
			"reflect"
		)

		type person struct {
			/* VarName VarType	`tags`
			 * the tags SHOULD BE separated by whitespace " " char
			*/
			Name 	string 	`type:"string" min:"2" max:"20"`
			Age		int 	`type:"numeric" min:"18" max:"99"`
			Country	string 	`type:"string" required:"true"`
		}

		func main() {
			// Filling 
			rafa := person{
				Name: "Rafael Breno",
				Age: 20,
				Country: "Brazil",
			}
			/* reflect.ValueOf(struct)
			 * return a Value "instance"
			 * "Value is the reflection interface to a Go value."
			*/
			v := reflect.ValueOf(rafa)

			// Return Value's Type
			typePerson := v.Type()
			// NumField() return the amount of fields
			for i := 0; i < v.NumField(); i++ {
				/* Field returns the field
				 * of that position
				 * 0 - Name
				 * 1 - Age
				 * 2 - Country
				*/
				field := typePerson.Field(i)

				/* field.Name will return the name of the field(Name, Age, Country)
				 * v.Field() will return it's value
				*/
				fmt.Println("Field: ", field.Name, " - Value: ", v.Field(i))
				/* field.Tag.Lookup will search a input and return 2 params
				 * tag - If found, return it's value 
				 * ok - boolean if the tag was found
				*/
				if tag, ok := field.Tag.Lookup("max"); ok {
					fmt.Println("Max: ", tag)
				}
				if tag, ok := field.Tag.Lookup("min"); ok {
					fmt.Println("Min: ", tag)
				}
				if tag, ok := field.Tag.Lookup("type"); ok {
					fmt.Println("Type: ", tag)
				}
				if tag, ok := field.Tag.Lookup("required"); ok {
					fmt.Println("Required: ", tag)
				}
				fmt.Println("-_-_-_-_-_-_-_-_-_-_-_-")
			}
		}
	```