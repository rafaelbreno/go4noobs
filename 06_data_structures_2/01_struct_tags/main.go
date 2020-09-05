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
		fmt.Println("Field: ", field.Name, " - Value: ", v.Field(i))
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