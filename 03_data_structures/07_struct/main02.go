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
	name string
	age  int

	/* Here is our custom type
	 * the date struct
	 */
	birthday date
}

func main() {
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
		name: "Rafael Breno",
		age:  20,
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
