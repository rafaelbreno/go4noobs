#### `enconding/json`
- JSON is a widely used data structure, you'll find it in most RESTful API and JS apps
- Useful links:
    - [_JSON Docs_](https://www.json.org/json-en.html) 
        - _"JSON (JavaScript Object Notation) is a lightweight data-interchange format. It is easy for humans to read and write. It is easy for machines to parse and generate."_
    - [_JSON MDN Docs_](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/JSON)
- Now back to the Go Package, I'll be following the [official doc](https://golang.org/pkg/encoding/json/)
- At the beginning it says to see [JSON and Go](https://golang.org/doc/articles/json_and_go.html)
##### First Exercise
- As a first exercise I will be creating a simple struct to user [`json.Marshal(v interface{})`](https://golang.org/pkg/encoding/json/#Marshal) function
- 	```go
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
	```
- One thing that's important to highlight here is the difference between attributes with the first leter being uppercase or lowercase
- The first try to use the `json.Marshal()` my struct was like this:
	- 	```go
			// Date struct
			type Date struct {
				day, month, year int
			}
			// Person struct
			type Person struct {
				Name     string
				Age      int
				Salary   float64
				Birthday Date
			}
		```
	- The output was: `{"Name":"John Doe","Age":20,"Salary":123.45,"Birthday":{}}`
	- As the documentation says:
		- _"The json package only accesses the exported fields of struct types (those that begin with an uppercase letter)."_
- This happened because: 
	- _"[Names](https://golang.org/doc/effective_go.html#names) are as important in Go as in any other language. They even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case."_