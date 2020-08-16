package main

import "fmt"

/* Creating a simple struct
 * will have 2 arguments
 * nameParts and completeName
 * completeName will be the concatenation of nameParts
 */
type person struct {
	nameParts    []string
	completeName string
}

/* Here we'll be receiving a Pointer
 * Thinks that a pointer, points to something
 * the case bellow
 * 'p' points to a variable of type 'person'
 */
func (p *person) concatName() {
	/* Because 'p' is a pointer of 'person'
	 * we can access and modify it's values
	 * without the need of returning something
	 */
	for _, namePart := range p.nameParts {
		// Accessing and modifying completeName
		p.completeName += namePart + " "
	}
}

func main() {
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
