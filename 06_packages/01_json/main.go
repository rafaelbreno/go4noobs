package main

// Importing package
import (
	// The package that we're used to
	"fmt"
	// Now importing the new package
	"encoding/json"
)

// Date struct
type Date struct {
	Day, Month, Year int
}

// Person struct
type Person struct {
	Name     string
	Age      int
	Salary   float64
	Birthday Date
}

func main() {
	p := Person{
		Name:   "John Doe",
		Age:    20,
		Salary: 123.45,
		Birthday: Date{
			Day:   14,
			Month: 6,
			Year:  1999,
		},
	}
	fmt.Println(p)
	/* Until here shouldn't be any issue
	 * Because structs were already studied at
	 * Data Structures / Structs
	 */

	b, err := json.Marshal(p)
	/* The json package only accesses the exported fields of struct types
	 * (those that begin with an uppercase letter).
	 */
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}
