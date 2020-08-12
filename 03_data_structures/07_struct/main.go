package main

import "fmt"

/*
 * For example we can create a struct of a person
 * and it will receive information that someone can have
 * Lie: name, height, adult(if he\she is an adult)
 */
var person struct {
	//	varName 	varType
	name   string
	height float64
	adult  bool
}

func main() {
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
