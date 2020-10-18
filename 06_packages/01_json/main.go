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

/* Seeing how Marshal func works */
func partOne() {
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

/* Seeing how Unmarshal works */
func partTwo() {
	var p Person

	b := []byte(`{"Name":"Ramon Doe","Age":32,"Salary":12345.67,"Birthday":{"Day":13,"Month":9,"Year":1987}}`)

	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}
	// {Ramon Doe 32 12345.67 {13 9 1987}}
	fmt.Println(p)
}

/* Marshal when there's tags that
 * doesn't exist in the original struct */
func partThree() {
	var p Person
	b := []byte(`{
					"Name":"Ramon Doe",
					"Age":32,
					"Random": "AANAuhUHSAKBSADA&!@$",
					"Salary":12345.67,
					"Birthday": {
						"Day":13,
						"Month":9,
						"Year":1987,
						"Random": "This won't show at the exported result"
					}
				}`)
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}
	/* {Ramon Doe 32 12345.67 {13 9 1987}}
	 * Unmarshal will decode only the fields that it can find in the destination type.
	 * That means the "random" fields will be "ignored"
	 */
	fmt.Println(p)
}

// Decoding without a struct
func partFour() {
	var i interface{}

	b := []byte(`{
					"Name":"Ramon Doe",
					"Random": "AANAuhUHSAKBSADA&!@$",
					"Birthday": {
						"Day":13,
						"Random": "This will show at the exported result"
					}
				}`)
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

func main() {
	partOne()

	partTwo()

	partThree()

	partFour()
}
